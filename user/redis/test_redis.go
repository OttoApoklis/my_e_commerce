package redis

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedis2() {
	// 连接到Redis服务器
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 根据实际情况修改
		Password: "",               // 如果没有密码则留空
		DB:       0,                // 使用默认DB
	})

	// 测试数据量
	const numKeys = 10000

	// 准备测试数据并写入Redis
	fmt.Println("开始写入数据到Redis...")
	start := time.Now()
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key:%d", i)
		value := fmt.Sprintf("value:%d", rand.Intn(1000000))
		err := rdb.Set(ctx, key, value, 0).Err()
		if err != nil {
			log.Printf("写入Redis失败: %v", err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("写入%d个键值对耗时: %s", numKeys, elapsed)

	// 随机查询性能测试
	fmt.Println("开始随机查询数据...")
	start = time.Now()
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key:%d", rand.Intn(numKeys))
		_, err := rdb.Get(ctx, key).Result()
		if err != nil {
			log.Printf("查询Redis失败: %v", err)
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("随机查询%d次耗时: %s", numKeys, elapsed)

	// 清理数据
	fmt.Println("开始清理数据...")
	start = time.Now()
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key:%d", i)
		rdb.Del(ctx, key)
	}
	elapsed = time.Since(start)
	fmt.Printf("删除%d个键值对耗时: %s", numKeys, elapsed)
}
