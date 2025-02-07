package utils

import (
	"context"
	"log"
	"my_e_commerce/config"
	"time"
)

func RedisSetStock(goodsID string, stock int, expiration time.Duration) {
	ctx := context.Background()
	rdb, err := config.GetRedisConnection()
	if err != nil {
		log.Printf("err : %+v", err)
	}
	ctx = context.Background()
	rdb.Set(ctx, goodsID, stock, expiration)
	return
}

func RedisSetKey(key string, expiration time.Duration) {
	ctx := context.Background()
	rdb, err := config.GetRedisConnection()
	if err != nil {
		log.Printf("err : %+v", err)
	}
	rdb.SetNX(ctx, key, nil, expiration)
	return
}
