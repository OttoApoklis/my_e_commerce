package Timer

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/enum"
	redis_test "my_e_commerce/redis_init"
	"my_e_commerce/service/serviceImpl"
	"my_e_commerce/utils"
	"time"
)

func ordertask() {

	ctx := context.Background()
	rdb, err := config.GetRedisConnection()
	if err != nil {
		log.Fatal("redis connection Error!")
	}
	opts := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    fmt.Sprintf("(%d", time.Now().Unix()),
		Offset: 0,
		Count:  0,
	}
	result, err := utils.FindToSortedSet(ctx, rdb, redis_test.ZSETKEY_ORDER, opts)
	orderService := serviceImpl.NewOrderServiceImpl()
	log.Printf("redis range result: %+v", result)
	for _, res := range result {
		// 循环更新订单状态
		orderNum, ok := res.Member.(string)
		if !ok {
			log.Printf("从redis获取的订单号，断言失败！")
			continue
		}
		log.Printf("orderNum: %v.", orderNum)
		ok, err = orderService.OrderStatusChange(orderNum, enum.SECKILL_ORDER_TIME_OUT)
		if err != nil {
			log.Printf("SeckillRecordStatusChange err caused by %v.", err)
		}
		if !ok {
			log.Printf("SeckillRecordStatusChange err")
		}
		// 更新成功就删除这个过期的订单键
		ok, err := utils.DeleteToSortedSet(ctx, rdb, redis_test.ZSETKEY_SECKILL_RECORD, orderNum)
		if err != nil {
			log.Printf("删除redis获取的订单号失败！ caused by:%v", err)
		}
		if !ok {
			log.Printf("删除redis获取的订单号失败！")
		}
	}
	result, err = utils.FindToSortedSet(ctx, rdb, redis_test.ZSETKEY_SECKILL_RECORD, opts)
	seckillRecordService := serviceImpl.NewSeckillServiceImpl()
	for _, res := range result {
		// 循环更新订单状态
		seckillRecordNum, ok := res.Member.(string)
		if !ok {
			log.Printf("从redis获取的秒杀号，断言失败！")
			continue
		}
		ok, err = seckillRecordService.SeckillRecordStatusChange(seckillRecordNum, enum.SECKILL_ORDER_TIME_OUT)
		if err != nil {
			log.Printf("SeckillRecordStatusChange err caused by %v.", err)
		}
		if !ok {
			log.Printf("SeckillRecordStatusChange err")
		}
		// 更新成功就删除这个过期的秒杀键
		ok, err := utils.DeleteToSortedSet(ctx, rdb, redis_test.ZSETKEY_SECKILL_RECORD, seckillRecordNum)
		if err != nil {
			log.Printf("删除redis获取的秒杀号失败！ caused by:%v", err)
		}
		if !ok {
			log.Printf("删除redis获取的秒杀号失败！")
		}
	}
	log.Printf("执行订单超时检查redis定时任务 %s", time.Now().Format(time.RFC1123))
	return
}

func OrderRedisTimer() {
	// ticker 用于定期发送事件， ticker.C 是一个只读的通道，每隔一段时间发送一个包含当前时间的对象，
	// 通过<-来接收该通道发送的对象可以达到延时触发事件的效果
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	ordertask()
	for {
		<-ticker.C
		ordertask()
	}
}
