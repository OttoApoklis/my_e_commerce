package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

var (
	conn               *amqp.Connection
	dsn                string
	rabbitmqOnce       sync.Once
	rabbitmqGetDSNOnce sync.Once
)

type RabbitmqConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"passwd"`
	Vhost    string `yaml:"vhost"`
}

type RabbitMQConfig struct {
	Rabbitmq RabbitmqConfig `yaml:"rabbitmq"`
}

func GetRabbitmqDSN() string {
	rabbitmqGetDSNOnce.Do(func() {
		var config RabbitMQConfig

		file, err := os.Open("config.yaml")
		if err != nil {
			log.Print(err)
		}
		defer func() {
			err := recover()
			if err != nil {
				log.Printf("err: %+v", err)
			}
		}()
		defer func() {
			file.Close()
		}()
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			log.Print(err)
		}
		dsn = fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
			config.Rabbitmq.User, config.Rabbitmq.Password, config.Rabbitmq.Host,
			config.Rabbitmq.Port, config.Rabbitmq.Vhost)
	})
	return dsn
}

func GetRabbitmqConnection() *amqp.Connection {
	var conn *amqp.Connection
	dsn = GetRabbitmqDSN()
	//rabbitmqOnce.Do(func() {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
	}
	log.Printf("mq connection is open")
	defer conn.Close()
	//})
	return conn
}
