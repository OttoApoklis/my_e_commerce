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

func newSeckillRecord(db *gorm.DB, opts ...gen.DOOption) seckillRecord {
	_seckillRecord := seckillRecord{}

	_seckillRecord.seckillRecordDo.UseDB(db, opts...)
	_seckillRecord.seckillRecordDo.UseModel(&model.SeckillRecord{})

	tableName := _seckillRecord.seckillRecordDo.TableName()
	_seckillRecord.ALL = field.NewAsterisk(tableName)
	_seckillRecord.ID = field.NewUint32(tableName, "id")
	_seckillRecord.UserID = field.NewUint32(tableName, "user_id")
	_seckillRecord.GoodsID = field.NewUint32(tableName, "goods_id")
	_seckillRecord.SecNum = field.NewString(tableName, "sec_num")
	_seckillRecord.OrderNum = field.NewString(tableName, "order_num")
	_seckillRecord.Price = field.NewUint32(tableName, "price")
	_seckillRecord.Status = field.NewUint32(tableName, "status")
	_seckillRecord.CreateTime = field.NewTime(tableName, "create_time")
	_seckillRecord.ModifyTime = field.NewTime(tableName, "modify_time")

	_seckillRecord.fillFieldMap()

	return _seckillRecord
}

// seckillRecord 秒杀记录表
type seckillRecord struct {
	seckillRecordDo

	ALL        field.Asterisk
	ID         field.Uint32 // ID
	UserID     field.Uint32 // 用户ID
	GoodsID    field.Uint32 // 商品ID
	SecNum     field.String // 秒杀号
	OrderNum   field.String // 订单号
	Price      field.Uint32 // 金额
	Status     field.Uint32 // 状态
	CreateTime field.Time   // 创建时间
	ModifyTime field.Time   // 修改时间

	fieldMap map[string]field.Expr
}

func (s seckillRecord) Table(newTableName string) *seckillRecord {
	s.seckillRecordDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s seckillRecord) As(alias string) *seckillRecord {
	s.seckillRecordDo.DO = *(s.seckillRecordDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *seckillRecord) updateTableName(table string) *seckillRecord {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint32(table, "id")
	s.UserID = field.NewUint32(table, "user_id")
	s.GoodsID = field.NewUint32(table, "goods_id")
	s.SecNum = field.NewString(table, "sec_num")
	s.OrderNum = field.NewString(table, "order_num")
	s.Price = field.NewUint32(table, "price")
	s.Status = field.NewUint32(table, "status")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")

	s.fillFieldMap()

	return s
}

func (s *seckillRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *seckillRecord) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["goods_id"] = s.GoodsID
	s.fieldMap["sec_num"] = s.SecNum
	s.fieldMap["order_num"] = s.OrderNum
	s.fieldMap["price"] = s.Price
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
}

func (s seckillRecord) clone(db *gorm.DB) seckillRecord {
	s.seckillRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s seckillRecord) replaceDB(db *gorm.DB) seckillRecord {
	s.seckillRecordDo.ReplaceDB(db)
	return s
}

type seckillRecordDo struct{ gen.DO }

type ISeckillRecordDo interface {
	gen.SubQuery
	Debug() ISeckillRecordDo
	WithContext(ctx context.Context) ISeckillRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISeckillRecordDo
	WriteDB() ISeckillRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISeckillRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISeckillRecordDo
	Not(conds ...gen.Condition) ISeckillRecordDo
	Or(conds ...gen.Condition) ISeckillRecordDo
	Select(conds ...field.Expr) ISeckillRecordDo
	Where(conds ...gen.Condition) ISeckillRecordDo
	Order(conds ...field.Expr) ISeckillRecordDo
	Distinct(cols ...field.Expr) ISeckillRecordDo
	Omit(cols ...field.Expr) ISeckillRecordDo
	Join(table schema.Tabler, on ...field.Expr) ISeckillRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISeckillRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISeckillRecordDo
	Group(cols ...field.Expr) ISeckillRecordDo
	Having(conds ...gen.Condition) ISeckillRecordDo
	Limit(limit int) ISeckillRecordDo
	Offset(offset int) ISeckillRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISeckillRecordDo
	Unscoped() ISeckillRecordDo
	Create(values ...*model.SeckillRecord) error
	CreateInBatches(values []*model.SeckillRecord, batchSize int) error
	Save(values ...*model.SeckillRecord) error
	First() (*model.SeckillRecord, error)
	Take() (*model.SeckillRecord, error)
	Last() (*model.SeckillRecord, error)
	Find() ([]*model.SeckillRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeckillRecord, err error)
	FindInBatches(result *[]*model.SeckillRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SeckillRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISeckillRecordDo
	Assign(attrs ...field.AssignExpr) ISeckillRecordDo
	Joins(fields ...field.RelationField) ISeckillRecordDo
	Preload(fields ...field.RelationField) ISeckillRecordDo
	FirstOrInit() (*model.SeckillRecord, error)
	FirstOrCreate() (*model.SeckillRecord, error)
	FindByPage(offset int, limit int) (result []*model.SeckillRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISeckillRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s seckillRecordDo) Debug() ISeckillRecordDo {
	return s.withDO(s.DO.Debug())
}

func (s seckillRecordDo) WithContext(ctx context.Context) ISeckillRecordDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s seckillRecordDo) ReadDB() ISeckillRecordDo {
	return s.Clauses(dbresolver.Read)
}

func (s seckillRecordDo) WriteDB() ISeckillRecordDo {
	return s.Clauses(dbresolver.Write)
}

func (s seckillRecordDo) Session(config *gorm.Session) ISeckillRecordDo {
	return s.withDO(s.DO.Session(config))
}

func (s seckillRecordDo) Clauses(conds ...clause.Expression) ISeckillRecordDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s seckillRecordDo) Returning(value interface{}, columns ...string) ISeckillRecordDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s seckillRecordDo) Not(conds ...gen.Condition) ISeckillRecordDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s seckillRecordDo) Or(conds ...gen.Condition) ISeckillRecordDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s seckillRecordDo) Select(conds ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s seckillRecordDo) Where(conds ...gen.Condition) ISeckillRecordDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s seckillRecordDo) Order(conds ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s seckillRecordDo) Distinct(cols ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s seckillRecordDo) Omit(cols ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s seckillRecordDo) Join(table schema.Tabler, on ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s seckillRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s seckillRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s seckillRecordDo) Group(cols ...field.Expr) ISeckillRecordDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s seckillRecordDo) Having(conds ...gen.Condition) ISeckillRecordDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s seckillRecordDo) Limit(limit int) ISeckillRecordDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s seckillRecordDo) Offset(offset int) ISeckillRecordDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s seckillRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISeckillRecordDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s seckillRecordDo) Unscoped() ISeckillRecordDo {
	return s.withDO(s.DO.Unscoped())
}

func (s seckillRecordDo) Create(values ...*model.SeckillRecord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s seckillRecordDo) CreateInBatches(values []*model.SeckillRecord, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s seckillRecordDo) Save(values ...*model.SeckillRecord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s seckillRecordDo) First() (*model.SeckillRecord, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeckillRecord), nil
	}
}

func (s seckillRecordDo) Take() (*model.SeckillRecord, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeckillRecord), nil
	}
}

func (s seckillRecordDo) Last() (*model.SeckillRecord, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeckillRecord), nil
	}
}

func (s seckillRecordDo) Find() ([]*model.SeckillRecord, error) {
	result, err := s.DO.Find()
	return result.([]*model.SeckillRecord), err
}

func (s seckillRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SeckillRecord, err error) {
	buf := make([]*model.SeckillRecord, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s seckillRecordDo) FindInBatches(result *[]*model.SeckillRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s seckillRecordDo) Attrs(attrs ...field.AssignExpr) ISeckillRecordDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s seckillRecordDo) Assign(attrs ...field.AssignExpr) ISeckillRecordDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s seckillRecordDo) Joins(fields ...field.RelationField) ISeckillRecordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s seckillRecordDo) Preload(fields ...field.RelationField) ISeckillRecordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s seckillRecordDo) FirstOrInit() (*model.SeckillRecord, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeckillRecord), nil
	}
}

func (s seckillRecordDo) FirstOrCreate() (*model.SeckillRecord, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SeckillRecord), nil
	}
}

func (s seckillRecordDo) FindByPage(offset int, limit int) (result []*model.SeckillRecord, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s seckillRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s seckillRecordDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s seckillRecordDo) Delete(models ...*model.SeckillRecord) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *seckillRecordDo) withDO(do gen.Dao) *seckillRecordDo {
	s.DO = *do.(*gen.DO)
	return s
}
