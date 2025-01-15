package redis_test

import (
	"golang.org/x/net/context"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	"strconv"
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
	for _, element := range seckillStocks {
		rdb.Set(ctx, strconv.FormatUint(uint64(*element.GoodsID), 10), element.Stock, 0)
	}
	return

}
