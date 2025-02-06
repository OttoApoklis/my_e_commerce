package Timer

import (
	"context"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	"strconv"
	"time"
)

func taskLog() {
	ctx := context.Background()
	rdb, err := config.GetRedisConnection()
	if err != nil {
		log.Fatal("redis connection Error!")
	}
	// 连接数据库加载库存信息
	db := config.GetDB()
	var seckillStocks []model.SeckillStock
	err = db.Model(&model.SeckillStock{}).Select("goods_id", "stock").
		Table("seckill_stock").Find(&seckillStocks).Error
	if err != nil {
		log.Printf("err : %+v", err)
	}
	ctx = context.Background()
	// 将秒杀库存的数据添加到数据库
	for _, element := range seckillStocks {
		go rdb.Set(ctx, strconv.FormatUint(uint64(*element.GoodsID), 10), element.Stock, 0)
	}
	log.Printf("执行redis定时任务 %s", time.Now().Format(time.RFC1123))
	return
}

func RedisTimer() {
	// ticker 用于定期发送事件， ticker.C 是一个只读的通道，每隔一段时间发送一个包含当前时间的对象，
	// 通过<-来接收该通道发送的对象可以达到延时触发事件的效果
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	taskLog()
	for {
		<-ticker.C
		taskLog()
	}
}
