package notificacion

import (
	"fmt"
	"log"
	"time"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/motordebusqueda"
	rabbit "github.com/rabbitmq/amqp091-go"
)

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
		"items", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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
		id := string(d.Body)
		log.Printf("Received a message: %s", d.Body)

		// Realizar acciones adicionales según el contenido del mensaje
		if id != "" {
			// Esperar 200ms antes de ejecutar Actualizar
			time.Sleep(200 * time.Millisecond)

			// Intentar actualizar y manejar el error
			err := actualizarConRetry(id)
			if err != nil {
				log.Printf("No se pudo corregir el ítem: %s", err.Error())
				// Continuar con el siguiente mensaje
				continue
			} else {
				log.Println(" - > Item actualizado")
			}
		}
		log.Println(" -> Escuchando mensajes... ")
	}
	return nil
}

// Función para intentar actualizar con retry
func actualizarConRetry(id string) error {
	// Intentar actualizar hasta 3 veces con un intervalo de 1 segundo entre intentos
	for i := 0; i < 3; i++ {
		err := motordebusqueda.Revisar(id)
		if err == nil {
			// Actualización exitosa
			return nil
		}
		log.Printf("Error al actualizar el ítem en Elasticsearch (intentando nuevamente): %s", err.Error())
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("Error al actualizar el ítem en Elasticsearch después de varios intentos")
}
