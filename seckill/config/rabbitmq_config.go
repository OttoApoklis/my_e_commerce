package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

// 连接池结构体
type RabbitMQConnectionPool struct {
	pool    chan *amqp.Connection
	factory func() (*amqp.Connection, error)
	size    int
	mu      sync.Mutex
}

var (
	dsn                string
	rabbitmqGetDSNOnce sync.Once
	rabbitmqPool       *RabbitMQConnectionPool
	rabbitmqPoolOnce   sync.Once
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

// 获取 RabbitMQ DSN
func GetRabbitmqDSN() string {
	rabbitmqGetDSNOnce.Do(func() {
		var config RabbitMQConfig

		file, err := os.Open("config.yaml")
		if err != nil {
			log.Print(err)
		}
		defer func() {
			if r := recover(); r != nil {
				log.Printf("err: %+v", r)
			}
		}()
		defer file.Close()

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

// 创建连接池
func NewRabbitMQConnectionPool(size int) (*RabbitMQConnectionPool, error) {
	dsn := GetRabbitmqDSN()
	factory := func() (*amqp.Connection, error) {
		return amqp.Dial(dsn)
	}
	pool := make(chan *amqp.Connection, size)
	for i := 0; i < size; i++ {
		conn, err := factory()
		if err != nil {
			return nil, err
		}
		pool <- conn
	}
	return &RabbitMQConnectionPool{
		pool:    pool,
		factory: factory,
		size:    size,
	}, nil
}

// 从连接池获取连接
func (p *RabbitMQConnectionPool) GetRabbitmqConn() (*amqp.Connection, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	select {
	case conn := <-p.pool:
		return conn, nil
	default:
		return p.factory()
	}
}

// 将连接归还到连接池
func (p *RabbitMQConnectionPool) PutRabbitmqConn(conn *amqp.Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.pool) < p.size {
		p.pool <- conn
	} else {
		conn.Close()
	}
}

// 获取 RabbitMQ 连接池
func GetRabbitmqConnectionPool(size int) (*RabbitMQConnectionPool, error) {
	var err error
	rabbitmqPoolOnce.Do(func() {
		rabbitmqPool, err = NewRabbitMQConnectionPool(size)
		if err != nil {
			log.Printf("Failed to create RabbitMQ connection pool: %s", err)
		}
	})
	return rabbitmqPool, err
}
