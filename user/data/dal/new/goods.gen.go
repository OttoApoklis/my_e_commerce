// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package new

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"my_e_commerce/data/dal/model"
)

func newGood(db *gorm.DB, opts ...gen.DOOption) good {
	_good := good{}

	_good.goodDo.UseDB(db, opts...)
	_good.goodDo.UseModel(&model.Good{})

	tableName := _good.goodDo.TableName()
	_good.ALL = field.NewAsterisk(tableName)
	_good.ID = field.NewUint32(tableName, "id")
	_good.GoodsNum = field.NewString(tableName, "goods_num")
	_good.GoodsName = field.NewString(tableName, "goods_name")
	_good.Price = field.NewFloat32(tableName, "price")
	_good.PicURL = field.NewString(tableName, "pic_url")
	_good.Seller = field.NewUint32(tableName, "seller")
	_good.CreateTime = field.NewTime(tableName, "create_time")
	_good.ModifyTime = field.NewTime(tableName, "modify_time")

	_good.fillFieldMap()

	return _good
}

// good 商品表
type good struct {
	goodDo

	ALL        field.Asterisk
	ID         field.Uint32
	GoodsNum   field.String  // 商品编号
	GoodsName  field.String  // 商品名字
	Price      field.Float32 // 价格
	PicURL     field.String  // 商品图片
	Seller     field.Uint32  // 卖家ID
	CreateTime field.Time    // 创建时间
	ModifyTime field.Time    // 修改时间

	fieldMap map[string]field.Expr
}

func (g good) Table(newTableName string) *good {
	g.goodDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g good) As(alias string) *good {
	g.goodDo.DO = *(g.goodDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *good) updateTableName(table string) *good {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewUint32(table, "id")
	g.GoodsNum = field.NewString(table, "goods_num")
	g.GoodsName = field.NewString(table, "goods_name")
	g.Price = field.NewFloat32(table, "price")
	g.PicURL = field.NewString(table, "pic_url")
	g.Seller = field.NewUint32(table, "seller")
	g.CreateTime = field.NewTime(table, "create_time")
	g.ModifyTime = field.NewTime(table, "modify_time")

	g.fillFieldMap()

	return g
}

func (g *good) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *good) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 8)
	g.fieldMap["id"] = g.ID
	g.fieldMap["goods_num"] = g.GoodsNum
	g.fieldMap["goods_name"] = g.GoodsName
	g.fieldMap["price"] = g.Price
	g.fieldMap["pic_url"] = g.PicURL
	g.fieldMap["seller"] = g.Seller
	g.fieldMap["create_time"] = g.CreateTime
	g.fieldMap["modify_time"] = g.ModifyTime
}

func (g good) clone(db *gorm.DB) good {
	g.goodDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g good) replaceDB(db *gorm.DB) good {
	g.goodDo.ReplaceDB(db)
	return g
}

type goodDo struct{ gen.DO }

type IGoodDo interface {
	gen.SubQuery
	Debug() IGoodDo
	WithContext(ctx context.Context) IGoodDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGoodDo
	WriteDB() IGoodDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGoodDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGoodDo
	Not(conds ...gen.Condition) IGoodDo
	Or(conds ...gen.Condition) IGoodDo
	Select(conds ...field.Expr) IGoodDo
	Where(conds ...gen.Condition) IGoodDo
	Order(conds ...field.Expr) IGoodDo
	Distinct(cols ...field.Expr) IGoodDo
	Omit(cols ...field.Expr) IGoodDo
	Join(table schema.Tabler, on ...field.Expr) IGoodDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGoodDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGoodDo
	Group(cols ...field.Expr) IGoodDo
	Having(conds ...gen.Condition) IGoodDo
	Limit(limit int) IGoodDo
	Offset(offset int) IGoodDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGoodDo
	Unscoped() IGoodDo
	Create(values ...*model.Good) error
	CreateInBatches(values []*model.Good, batchSize int) error
	Save(values ...*model.Good) error
	First() (*model.Good, error)
	Take() (*model.Good, error)
	Last() (*model.Good, error)
	Find() ([]*model.Good, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Good, err error)
	FindInBatches(result *[]*model.Good, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Good) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGoodDo
	Assign(attrs ...field.AssignExpr) IGoodDo
	Joins(fields ...field.RelationField) IGoodDo
	Preload(fields ...field.RelationField) IGoodDo
	FirstOrInit() (*model.Good, error)
	FirstOrCreate() (*model.Good, error)
	FindByPage(offset int, limit int) (result []*model.Good, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGoodDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g goodDo) Debug() IGoodDo {
	return g.withDO(g.DO.Debug())
}

func (g goodDo) WithContext(ctx context.Context) IGoodDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goodDo) ReadDB() IGoodDo {
	return g.Clauses(dbresolver.Read)
}

func (g goodDo) WriteDB() IGoodDo {
	return g.Clauses(dbresolver.Write)
}

func (g goodDo) Session(config *gorm.Session) IGoodDo {
	return g.withDO(g.DO.Session(config))
}

func (g goodDo) Clauses(conds ...clause.Expression) IGoodDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goodDo) Returning(value interface{}, columns ...string) IGoodDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goodDo) Not(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goodDo) Or(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goodDo) Select(conds ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goodDo) Where(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goodDo) Order(conds ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goodDo) Distinct(cols ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goodDo) Omit(cols ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goodDo) Join(table schema.Tabler, on ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goodDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGoodDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goodDo) RightJoin(table schema.Tabler, on ...field.Expr) IGoodDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goodDo) Group(cols ...field.Expr) IGoodDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goodDo) Having(conds ...gen.Condition) IGoodDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goodDo) Limit(limit int) IGoodDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goodDo) Offset(offset int) IGoodDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goodDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGoodDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goodDo) Unscoped() IGoodDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goodDo) Create(values ...*model.Good) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goodDo) CreateInBatches(values []*model.Good, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goodDo) Save(values ...*model.Good) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goodDo) First() (*model.Good, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) Take() (*model.Good, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) Last() (*model.Good, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) Find() ([]*model.Good, error) {
	result, err := g.DO.Find()
	return result.([]*model.Good), err
}

func (g goodDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Good, err error) {
	buf := make([]*model.Good, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goodDo) FindInBatches(result *[]*model.Good, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goodDo) Attrs(attrs ...field.AssignExpr) IGoodDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goodDo) Assign(attrs ...field.AssignExpr) IGoodDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goodDo) Joins(fields ...field.RelationField) IGoodDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goodDo) Preload(fields ...field.RelationField) IGoodDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goodDo) FirstOrInit() (*model.Good, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) FirstOrCreate() (*model.Good, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Good), nil
	}
}

func (g goodDo) FindByPage(offset int, limit int) (result []*model.Good, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g goodDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goodDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goodDo) Delete(models ...*model.Good) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goodDo) withDO(do gen.Dao) *goodDo {
	g.DO = *do.(*gen.DO)
	return g
}
