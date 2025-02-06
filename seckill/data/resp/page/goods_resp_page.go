package page

import "my_e_commerce/data/dal/model"

type GoodsRespPage struct {
	PageSize   uint32
	PageNumber uint32
	TotalPages uint32
	Data       []model.Good
}
