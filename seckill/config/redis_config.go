package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	redisOnce   sync.Once
	redisClient *redis.Client
	redisErr    error
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"passwd"`
	DB       int    `yaml:"db"` // 使用 int 类型表示数据库编号
}

type RConfig struct {
	Redis RedisConfig `yaml:"redis"`
}

// GetRedisClient 初始化并返回 Redis 客户端实例（单例模式）
func GetRedisClient() (*redis.Client, error) {
	redisOnce.Do(func() {
		var config RConfig
		if err := viper.Unmarshal(&config); err != nil {
			redisErr = fmt.Errorf("unable to decode into struct: %v", err)
			return
		}

		// 初始化 Redis 客户端
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
			Username: config.Redis.User,
			Password: config.Redis.Password,
			DB:       config.Redis.DB,
			PoolSize: 20, // 设置连接池大小
		})

		// 测试连接
		pong, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			redisErr = fmt.Errorf("failed to connect to Redis: %v", err)
			return
		}
		log.Printf("Redis connected: %s", pong)
	})

	return redisClient, redisErr
}

// GetRedisConnection 返回 Redis 客户端实例
func GetRedisConnection() (*redis.Client, error) {
	return GetRedisClient()
}
