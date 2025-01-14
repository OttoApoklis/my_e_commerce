package redis_test

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"my_e_commerce/config"
	"testing"
	"time"
)

func TestRedis2(t *testing.T) {
	ctx := context.Background()
	rdb, err := config.GetRedisConnection()
	if err != nil {
		log.Fatal("redis connection Error!")
	}
	rdb.Set(ctx, "key", "value", 10*time.Second)
	value := rdb.Get(ctx, "key")
	fmt.Printf("redis value: %+v", value)
}
