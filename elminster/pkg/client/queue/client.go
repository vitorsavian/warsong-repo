package queue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
}

func CreateClient() *Client {
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Pub() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
}
