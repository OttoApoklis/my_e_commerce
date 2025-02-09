package serviceImpl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"my_e_commerce/config"
	"my_e_commerce/data/dal/model"
	model2 "my_e_commerce/data/req"
	"my_e_commerce/data/req/page"
	"my_e_commerce/enum"
	"strconv"
	"strings"
)

type GoodsServiceImpl struct {
}

func NewGoodsServiceImpl() *GoodsServiceImpl { return &GoodsServiceImpl{} }

func (s *GoodsServiceImpl) CreateGoods(db *gorm.DB, goods *model.Good) error {
	var goodss []model.Good
	if err := db.Where("goods_num", goods.GoodsNum).Find(&goodss).Error; err != nil {
		log.Printf("err from quota find in createGoods: %+v", err)
		return err
	}
	if goodss != nil && len(goodss) != 0 {
		log.Printf("err from goods create because of repeateable")
		return errors.New("repeatable in goods create")
	}
	err := db.Save(goods).Error
	if err != nil {
		log.Printf("error from Goods create: %+v", err)
		return err
	}
	return nil
}

func (s *GoodsServiceImpl) GetGoods(goodsID uint32) ([]model.Good, error) {
	db := config.GetDB()
	goods := []model.Good{}
	db = db.Select("id", "goods_num", "goods_name", "price",
		"pic_url", "seller")
	err := db.Where("id = ?", goodsID).Find(&goods).Error
	log.Printf("%+v", goods)
	if err != nil {
		log.Printf(" err: %+v", err)
		return nil, err
	}
	if goods == nil {
		log.Printf("Get Goods is nil")
	}
	log.Printf("%+v", goods)
	return goods, nil
}

func (s *GoodsServiceImpl) GetGoodsByUser(goodsReq model2.GoodsGetUserReq) ([]model.Good, error) {
	db := config.GetDB()
	goods := []model.Good{}
	db = db.Select("id", "goods_num", "goods_name", "price",
		"pic_url", "seller")
	if goodsReq.GoodsType != nil {
		numPrefix, err := enum.GetNumPrefix(*goodsReq.GoodsType)
		if err != nil {
			log.Printf("getNumPrefix err caused by %v.", err)
			return goods, err
		}
		db = db.Where("goods_num like ?", fmt.Sprintf("%s%%", numPrefix))
	}
	if goodsReq.GoodsName != nil {
		db = db.Where("goods_name = ?", goodsReq.GoodsName)
	}
	err := db.Find(&goods).Error
	log.Printf("%+v", goods)
	if err != nil {
		log.Printf(" err: %+v", err)
		return nil, err
	}
	if goods == nil {
		log.Printf("Get Goods is nil")
	}
	log.Printf("%+v", goods)
	return goods, nil
}

func (s *GoodsServiceImpl) GetGoodsInPage(req model2.GoodsGetPageReq) (page.GoodsRespPage, error) {
	db1 := config.GetDB()
	var (
		offset uint32
		size   uint32
	)
	var goodsRespPage page.GoodsRespPage

	if req.PageSize == 0 {
		size = 10
	} else {
		size = req.PageSize
	}
	if req.PageNum == 0 {
		offset = 0
		goodsRespPage.PageNumber = 1
	} else {
		goodsRespPage.PageNumber = offset + 1
		offset = (req.PageNum - 1) * size
		log.Printf("offset %d", offset)
	}
	sql := strings.Builder{}
	sql.WriteString("select count(*) from goods ")
	var where1 *string
	var where2 *string
	if req.GoodsType != nil {
		prefix, err := enum.GetNumPrefix(*req.GoodsType)
		if err != nil {
			log.Printf("GetNumPrefix err caused by %v", err)
		}
		where1 = &prefix
	}
	if req.GoodsName != nil {
		str := "goods_name like \"" + *req.GoodsName + "%\""
		where2 = &str
	}
	if where1 != nil || where2 != nil {
		sql.WriteString("where ")
		if where1 != nil {
			sql.WriteString(*where1)
		}
		if where2 != nil {
			sql.WriteString(*where2)
		}
	}
	log.Printf("sql %s", sql.String())
	goods := []model.Good{}

	var total uint32
	err := db1.Raw(sql.String()).Scan(&total).Error
	if err != nil {
		log.Printf("err from GetGoodsInPage select total, err: %+v", err)
	}
	goodsRespPage.TotalPages = total / size
	if total%size > 0 {
		goodsRespPage.TotalPages++
	}
	goodsRespPage.PageSize = size

	log.Printf("total: %d", total)
	if offset != 0 && total < offset {
		log.Printf("no data in this page from GetGoodsInPage")
		return goodsRespPage, nil
	}
	db2 := config.GetDB()
	sql2 := strings.Builder{}
	sql2.WriteString("select * from goods ")
	db2 = db2.Table("goods")
	if where1 != nil || where2 != nil {
		sql2.WriteString("where ")
		if where1 != nil {
			sql2.WriteString(*where1)
		}
		if where2 != nil {
			sql2.WriteString(*where2)
		}
	}

	log.Printf("sql2 %s", sql2.String())
	log.Printf("offset")
	log.Println(offset)
	log.Println(int(offset))
	//if offset != 0 {
	//	sql2.WriteString(" offset " + strconv.Itoa(int(offset)))
	//}
	//if size != 0 {
	//	sql2.WriteString(" limit " + strconv.Itoa(int(size)))
	//}
	sql2.WriteString(" limit " + strconv.Itoa(int(offset)) + "," + strconv.Itoa(int(size)))

	err = db2.Raw(sql2.String()).Find(&goods).Error

	if err != nil {
		log.Printf("err from GetGoodsInPage, err: %+v", err)
	}
	goodsRespPage.Data = goods
	return goodsRespPage, nil
}

func (s *GoodsServiceImpl) UpdateGoods(db *gorm.DB, req *model2.GoodsUpdateReq) error {
	var goods model2.GoodsUpdateReq
	goods = *req
	dbMessage := db.Model(&model.Good{}).
		Where("goods_num = ?", *goods.GoodsNum).
		Limit(1).
		Update("price", goods.Price)
	//if dbMessage.RowsAffected == 0 {
	//	return errors.New("查不到该数据")
	//}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Goods update: %+v", err)
		return err
	}
	return nil
}

func (s *GoodsServiceImpl) DeleteGoodsById(db *gorm.DB, ID uint32) error {
	dbMessage := db.Where("id = ?", ID).Delete(&model.Good{})
	if dbMessage.RowsAffected == 0 {
		log.Printf("rows affected is zero in DeleteGoodsById")
		return errors.New("查不到数据")
	}
	err := dbMessage.Error
	if err != nil {
		log.Printf("error from Goods delete: %+v", err)
	}
	return nil
}
