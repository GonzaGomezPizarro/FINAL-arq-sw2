package notificacion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"io/ioutil"
	"net/http"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/items/dto"
	rabbit "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Send(id string) {
	conn, err := rabbit.Dial("amqp://guest:guest@rabbit:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"items", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := id
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		rabbit.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func Receive() error {
	conn, err := rabbit.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Trabajo_para_items", // name
		false,                // no durable
		true,
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	replyQueue, err := ch.QueueDeclare(
		"Respuestas_para_items",
		false, // no durable
		false, // not delete when unused
		false, // no exclusivo
		false, // no-wait
		nil,   // argumentos
	)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}
	log.Println(" -> Escuchando mensajes... ")

	// Consumir mensajes continuamente
	for d := range msgs {
		log.Printf(" [x] Received %s\n", d.Body)

		// Parsear el mensaje recibido
		var trabajo dto.TrabajoItem
		if err := json.Unmarshal(d.Body, &trabajo); err != nil {
			log.Println("Failed to parse message:", err)
			return err
		}

		metodo := trabajo.Metodo
		url := trabajo.Url
		jsonn := trabajo.Jsonn

		// Realizar solicitudes internas basadas en el método y la URL
		items, httpStatusCode, err := solicitudInterna(metodo, url, jsonn)
		if err != nil {
			log.Println("Internal request failed:", err)
			return err
		}

		// Construir la respuesta en formato dto.RespuestaItem
		respuesta := dto.RespuestaItem{
			Items:          items,
			HttpStatusCode: httpStatusCode,
		}

		// Serializar la respuesta a JSON
		jsonResponse, err := json.Marshal(respuesta)
		if err != nil {
			log.Println("Failed to marshal response to JSON:", err)
			return err
		}

		// Enviar la respuesta a la cola de respuesta
		err = ch.PublishWithContext(context.Background(), "", replyQueue.Name, false, false, rabbit.Publishing{
			ContentType:   "application/json",
			CorrelationId: d.CorrelationId,
			Body:          jsonResponse,
		})

		if err != nil {
			log.Println("Failed to send response to reply queue:", err)
			return err
		}
		log.Printf(" [x] Sent response to reply queue\n")
	}

	return nil
}

func solicitudInterna(metodo string, path string, jsonn string) (dto.Items, int, error) {
	url := "http://localhost:8091" + path //veremosssss
	var items dto.Items

	// Convertir el JSON en bytes, solo si no está vacío
	var jsonBytes []byte
	if jsonn != "" {
		jsonBytes = []byte(jsonn)
	}

	// Crear una nueva solicitud HTTP
	var req *http.Request
	var err error
	if metodo == "GET" || metodo == "DELETE" {
		req, err = http.NewRequest(metodo, url, nil)
	} else {
		req, err = http.NewRequest(metodo, url, bytes.NewBuffer(jsonBytes))
	}

	if err != nil {
		return dto.Items{}, 500, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Establecer la cabecera Content-Type si el JSON no está vacío
	if jsonn != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Realizar la solicitud HTTP interna
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return dto.Items{}, 500, fmt.Errorf("failed to perform internal HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dto.Items{}, 500, fmt.Errorf("failed to read response body: %v", err)
	}

	// Deserializar la respuesta JSON en la estructura dto.Items si la respuesta no está vacía
	if len(body) > 0 {
		if path == "/items" {
			if err := json.Unmarshal(body, &items); err != nil {
				return items, 0, fmt.Errorf("failed to unmarshal response body: %v", err)
			}
		} else {
			var item dto.Item
			if err := json.Unmarshal(body, &item); err != nil {
				return items, 0, fmt.Errorf("failed to unmarshal response body: %v", err)
			}
			items = append(items, item)
		}

	}

	return items, resp.StatusCode, nil
}
