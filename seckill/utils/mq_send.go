package utils

import (
	"encoding/json"
	"log"
	"my_e_commerce/config"
	model "my_e_commerce/data/req"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

// 发送函数
func SendMessageOnce(queueName string, data model.SeckillReq) {
	connPool, err := config.GetRabbitmqConnectionPool(100)
	// 获取连接
	conn, err := connPool.GetRabbitmqConn()
	defer func() {
		// 回收连接
		connPool.PutRabbitmqConn(conn)
	}()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body, err := json.Marshal(data)
	failOnError(err, "Failed to marshal data")

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
