package service

import model "my_e_commerce/data/req"

type SeckillRecordService interface {
	CreateSecRecord(service *SeckillRecordService) error
	UpdateSeckillRecord(seckillRecordReq *model.SeckillRecordReq) error
	GetSeckillRecord(service *SeckillRecordService) ([]*SeckillRecordService, error)
}
