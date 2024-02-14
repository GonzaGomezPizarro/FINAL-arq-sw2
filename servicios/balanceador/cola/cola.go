package cola

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/balanceador/dto"
	rabbit "github.com/rabbitmq/amqp091-go"
)

var idMensaje int = 0

func CheckConection() error {
	_, err := rabbit.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		return err
	}
	return nil
}

func SendToItems(metodo string, url string, jsonn string) (dto.Items, int, error) {
	conn, err := rabbit.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return dto.Items{}, 503, err // Servicio no disponible
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err)
		return dto.Items{}, 503, err // Servicio no disponible
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Trabajo_para_items", // name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue:", err)
		return dto.Items{}, 503, err // Servicio no disponible
	}

	replyQueue, err := ch.QueueDeclare(
		"Respuestas_para_items",
		false, // no durable
		true,  // auto eliminación cuando no hay consumidores
		false, // no exclusivo
		false, // no-wait
		nil,   // argumentos
	)
	if err != nil {
		log.Println("Failed to declare a response queue:", err)
		return dto.Items{}, 503, err // Servicio no disponible
	}

	msgs, err := ch.Consume(
		replyQueue.Name, // nombre de la cola
		"",              // consumer tag
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:", err)
		return dto.Items{}, 503, err // Servicio no disponible
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idMensaje++
	props := rabbit.Publishing{
		ContentType:   "text/plain",
		CorrelationId: strconv.Itoa(idMensaje),
		ReplyTo:       replyQueue.Name,
	}

	body := metodo + "." + url + "." + jsonn
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, props)
	if err != nil {
		log.Println("Failed to publish a message:", err)
		return dto.Items{}, 503, err // Servicio no disponible
	}
	log.Printf(" [x] Sent %s\n", body)

	select {
	case <-ctx.Done():
		log.Println("Timeout waiting for response")
		return dto.Items{}, 504, err // Gateway Timeout
	case d := <-msgs:
		if d.CorrelationId == strconv.Itoa(idMensaje) {
			log.Printf(" [x] Response: %s\n", d.Body)
			var response dto.RespuestaItem
			if err := json.Unmarshal(d.Body, &response); err != nil {
				log.Println("Error al deserializar la respuesta JSON:", err)
				return dto.Items{}, 500, err // Error interno del servidor
			}
			return response.Items, response.HttpStatusCode, nil
		} else {
			err = ch.Publish("", replyQueue.Name, false, false, rabbit.Publishing{
				ContentType:   "text/plain",
				CorrelationId: d.CorrelationId,
				Body:          d.Body,
			})
			if err != nil {
				log.Println("Failed to republish a message:", err)
				return dto.Items{}, 503, err // Servicio no disponible
			}
			log.Printf(" [x] Republished message with incorrect CorrelationId: %s\n", d.CorrelationId)
		}
	}

	//aca nunca llega...
	return nil, 500, err
}
