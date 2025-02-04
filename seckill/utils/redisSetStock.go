package utils

import (
	"context"
	"log"
	"my_e_commerce/config"
)

func RedisSetStock(goodsID string, stock int) {
	ctx := context.Background()
	rdb, err := config.GetRedisConnection()
	if err != nil {
		log.Fatal("redis connection Error!")
	}
	if err != nil {
		log.Printf("err : %+v", err)
	}
	ctx = context.Background()
	rdb.Set(ctx, goodsID, stock, 0)
	return
}
