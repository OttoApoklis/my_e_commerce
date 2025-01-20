package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
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

// RedisInstance 包含 redis.Client 实例
type RedisInstance struct {
	client *redis.Client
}

// GetRedisClient 初始化并返回 Redis 客户端实例
func GetRedisClient(config RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username: config.User,
		Password: config.Password,
		DB:       config.DB,
	})

	// 测试连接
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	log.Printf("Redis connected: %s", pong)
	return rdb, nil
}

func init() {
	// 设置 Viper 解析 YAML 配置文件
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
}

func GetRedisConnection() (*redis.Client, error) {
	var config RConfig
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}
	rdb, err := GetRedisClient(config.Redis)
	if err != nil {
		log.Printf("Failed to initialize Redis client: %v", err)
	}
	return rdb, err
}

func example() {
	var config RConfig
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}
	rdb, err := GetRedisClient(config.Redis)
	if err != nil {
		log.Printf("Failed to initialize Redis client: %v", err)
	}
	ctx := context.Background()
	// 使用 rdb 进行操作
	// 示例：设置一个键值对
	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Printf("Failed to set key: %v", err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Printf("Failed to get key: %v", err)
	}
	fmt.Println("Key value:", val)
}
