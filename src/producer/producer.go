package producer

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func InitProducer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	"demo.queue", // name
	// 	false,        // durable
	// 	false,        // delete when unused
	// 	false,        // exclusive
	// 	false,        // no-wait
	// 	nil,          // arguments
	// )
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var jsonBody string

	for i := 1; i < 100000; i++ {
		jsonBody = fmt.Sprintf("{'" + strconv.Itoa(i) + "': 'test'}")
		err = ch.PublishWithContext(ctx,
			"demo.exchange", // exchange
			"",              // routing key
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(jsonBody),
			})
		if i%100 == 0 {
			time.Sleep(2 * time.Second)
		}
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", jsonBody)
	}
	for true {
		log.Printf(" sleep ")
		time.Sleep(60 * time.Second)

	}
}
