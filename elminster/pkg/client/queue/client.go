package queue

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
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

func CreateConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		logrus.Errorf("Failed to create connection with rabbitmq: %v", err)
		return nil, err
	}

	return conn, nil
}
