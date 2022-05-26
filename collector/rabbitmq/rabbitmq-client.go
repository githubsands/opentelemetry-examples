package main

import (
	"os"

	"github.com/streadway/amqp"
)

func main() {
	_ = os.Getenv("RABBITMQ_ENDPOINT")
	_ = os.Getenv("RABBITMQ_PASSWORD")
	_ = os.Getenv("RABBITMQ_USER")

	conn, err := amqp.Dial("amqp://user:userpass@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		panic(err)
	}

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		panic(err)
	}

	for {
	}
}
