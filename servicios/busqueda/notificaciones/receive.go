package notificacion

import (
	"log"

	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/busqueda/elasticsearch"
	rabbit "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Receive(messages chan<- string) {
	conn, err := rabbit.Dial("amqp://guest:guest@localhost:5672/")
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Consumir mensajes continuamente
	for d := range msgs {
		id := string(d.Body)
		log.Printf("Received a message: %s", d.Body)
		// Realizar acciones adicionales según el contenido del mensaje
		if id != "" {
			err := elasticsearch.Actualizar(id)
			if err != nil {
				log.Println(err.Error())
			}
		}
		messages <- string(d.Body)
		// Agregar más lógica aquí según lo que desees hacer con el mensaje recibido
	}
}
