package rabbitmq

import (
	"encoding/json"
	"log"
	model "my_e_commerce/data/resp"
	"time"

	"github.com/streadway/amqp"
	"golang.org/x/time/rate"
)

// 发送函数
func sendMessage(conn *amqp.Connection, queueName string, data model.SeckillReq, frequency time.Duration) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body, err := json.Marshal(data)
	failOnError(err, "Failed to marshal data")

	ticker := time.NewTicker(frequency)
	defer ticker.Stop()

	for range ticker.C {
		err = ch.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		failOnError(err, "Failed to publish a message")
		log.Printf("Sent: %+v", data)
	}
}

// 接收函数（带限流）
func receiveMessage(conn *amqp.Connection, queueName string, limit rate.Limit) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no - wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto - ack
		false,  // exclusive
		false,  // no - local
		false,  // no - wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	limiter := rate.NewLimiter(limit, 1)

	for d := range msgs {
		if limiter.Allow() {
			var data model.SeckillReq
			err = json.Unmarshal(d.Body, &data)
			failOnError(err, "Failed to unmarshal data")
			log.Printf("Received: %+v", data)
		} else {
			log.Println("Rate limit exceeded, skipping message")
		}
	}
}

func RabbitmqTest2() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	queueName := "test_queue"

	var data model.SeckillReq
	data.GoodsID = 1
	data.GoodsAmount = 1
	go sendMessage(conn, queueName, data, 1*time.Second)

	receiveMessage(conn, queueName, rate.Every(2*time.Millisecond))
}
