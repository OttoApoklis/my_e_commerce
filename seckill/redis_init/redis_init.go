package redis_test

import (
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	"strconv"

	"golang.org/x/net/context"
)

func RedisInit() {
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
	// 清空redis数据库原有信息防止干扰
	rdb.FlushDB(ctx).Err()
	if err != nil {
		log.Printf("清空redis数据库失败！")
	}
	// 将秒杀库存的数据添加到数据库
	for _, element := range seckillStocks {
		rdb.Set(ctx, strconv.FormatUint(uint64(*element.GoodsID), 10), element.Stock, 0)
	}
	return

}
