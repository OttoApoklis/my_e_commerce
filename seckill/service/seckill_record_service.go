package service

import (
	model2 "my_e_commerce/data/dal/model"
	model "my_e_commerce/data/req"
)

type SeckillRecordService interface {
	CreateSeckillRecord(seckillRecordReq *model.SeckillRecordReq) error
	UpdateSeckillRecord(seckillRecordReq *model.SeckillRecordReq) error
	GetSeckillRecordByID(id uint32) ([]model2.SeckillRecord, error)
	GetSeckillRecord(seckillNum string) ([]*model2.SeckillRecord, error)
	GetSeckillRecordByUser(seckillGetRecord model.SeckillRecordGetReq) ([]*model2.SeckillRecord, error)
	GetSeckillRecordBySecNum(seckillGetRecord model.SeckillRecordBySecNumReq) ([]*model2.SeckillRecord, error)
	GetSeckillRecordByUserLast(seckillGetRecord model.SeckillRecordGetReq) ([]*model2.SeckillRecord, error)
	SeckillRecordStatusChange(seckillRecordNum string, status uint32) (bool, error)
	DeleteRedisSeckillRecord(seckillRecordNum string) (bool, error)
}
