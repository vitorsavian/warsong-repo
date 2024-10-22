package sub

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

func Run(msgs <-chan amqp.Delivery) bool {
	logrus.Infof("Waiting for messages.")

	for d := range msgs {
		logrus.Printf("Received a message: %s", d.Body)
	}

	return true
}

func Pub() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		logrus.Errorf("Failed to create connection with rabbitmq: %v", err)
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logrus.Errorf("Failed to listen to rabbitmq channel: %v", err)
		return err
	}

	defer ch.Close()
	var queue amqp.Queue
	if queue, err = ch.QueueDeclare(
		"character",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		logrus.Errorf("Failed to declare queue: %v", err)
		return err
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		logrus.Errorf("Failed to create channel for sub: %v", err)
		return err
	}

	for Run(msgs) {

	}

	logrus.Info("Retrying connection")

	return Pub()
}
