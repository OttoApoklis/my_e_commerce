package config

import (
	"github.com/streadway/amqp"
	"log"
)

func GetRabbitmqConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	log.Printf("mq connection is open")
	defer conn.Close()
	return conn
}
