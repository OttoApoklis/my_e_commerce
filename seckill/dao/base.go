package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"my_e_commerce/data/response"
	"trpc.group/trpc-go/trpc-go/log"
)

type contextKey int

var dbKey contextKey = 0

// Base 定义基础DB结构
type Base struct {
	db *gorm.DB
}

// 初始化DB
func NewDB(db *gorm.DB) *Base {
	return &Base{
		db: db,
	}
}

// withDB
func withDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

// GetDB 获取DB变量
func (b *Base) GetDB(ctx context.Context) *gorm.DB {
	v := ctx.Value(dbKey)
	if v == nil {
		return b.db
	}
	db := v.(*gorm.DB)
	return db
}

// TransactionDo 事务执行方法
func (b *Base) TransactionDo(ctx context.Context, fn func(context.Context) error) error {
	var err error
	db := b.db.Begin()
	if err = db.Error; err != nil {
		return err
	}

	defer func() {
		if e := recover(); e != nil {
			log.Info("start to rollback")
			db.Rollback()
			panic(e)
		} else if err != nil {
			log.Info("start to rollback")
			db.Rollback()
		} else {
			err = db.Commit().Error
			if err != nil {
				db.Rollback()
			}
		}
	}()

	ctx = withDB(ctx, db)
	err = fn(ctx)
	return err
}

// Update 更新操作
func (b *Base) Update(ctx context.Context, table string,
	conditions map[string][]interface{}, data map[string]interface{}) (int64, error) {
	db := b.GetDB(ctx)

	query := db.Table(table)
	for key, values := range conditions {
		query = query.Where(key, values...)
	}

	query = query.Updates(data)
	rowsAffected, err := query.RowsAffected, query.Error
	if err != nil {
		log.Errorf("update failed(%v)", err)
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateForLimit 更新操作
func (b *Base) UpdateForLimit(ctx context.Context, table string,
	conditions map[string][]interface{}, data map[string]interface{},
	rows int64) (int64, error) {
	db := b.GetDB(ctx)
	query := db.Table(table)
	for key, values := range conditions {
		query = query.Where(key, values...)
	}
	query = query.Limit(rows).Updates(data)
	rowsAffected, err := query.RowsAffected, query.Error
	if err != nil {
		log.Errorf("update failed(%v)", err)
		return 0, err
	}

	return rowsAffected, nil
}

// Delete 删除操作
func (b *Base) Delete(ctx context.Context, table string, model interface{},
	conditions map[string][]interface{}) (int64, error) {
	db := b.GetDB(ctx)
	query := db.Table(table)
	for key, values := range conditions {
		query = query.Where(key, values...)
	}
	query = query.Delete(model)
	rowsAffected, err := query.RowsAffected, query.Error

	if err != nil {
		log.Errorf("delete failed(%v)", err)
		return 0, err
	}

	return rowsAffected, nil
}

// Count 查询总记录数
func (b *Base) Count(ctx context.Context, table string, conditions map[string][]interface{}) (int64, error) {
	var totalNum int64
	db := b.GetDB(ctx)
	query := db.Table(table)

	for key, values := range conditions {
		query = query.Where(key, values...)
	}

	err := query.Count(&totalNum).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Info("record is not found")
		return 0, nil
	} else if err != nil {
		log.Errorf("get record failed(%v)", err)
	}

	return totalNum, err
}

// Create 创建操作
func (b *Base) Create(ctx context.Context, table string, model interface{}) error {
	db := b.GetDB(ctx)
	if err := db.Table(table).
		Create(model).
		Scan(model).
		Error; err != nil {
		log.Errorf("create failed(%v)", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.New(response.GetErrMsg(response.ERR_DB_DBDuplicateERROR))
		}
		return err
	}

	return nil
}

// Save 创建操作
func (b *Base) Save(ctx context.Context, table string, model interface{}) error {
	db := b.GetDB(ctx)
	if err := db.Table(table).
		Save(model).
		Error; err != nil {
		log.Errorf("save failed(%v)", err)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.New(response.GetErrMsg(response.ERR_DB_DBDuplicateERROR))
		}
		return err
	}

	return nil
}

// GetByPage 分页查询操作
func (b *Base) GetByPage(ctx context.Context, table string, res interface{}, conditions map[string][]interface{},
	order string, rows int64, offset int64) error {
	db := b.GetDB(ctx)
	query := db.Table(table)

	for key, values := range conditions {
		query = query.Where(key, values...)
	}

	if len(order) > 0 {
		query = query.Order(order)
	}

	err := query.Limit(rows).Offset(offset).Find(res).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Info("record is not found")
		return nil
	} else if err != nil {
		log.Errorf("get failed(%v)", err)
	}

	return err
}

// GetByPageAndCount 分页查询操作包含count
func (b *Base) GetByPageAndCount(ctx context.Context, table string, res interface{}, count *int,
	conditions map[string][]interface{}, order string, rows int64, offset int64) error {
	db := b.GetDB(ctx)
	query := db.Table(table)

	for key, values := range conditions {
		query = query.Where(key, values...)
	}

	if len(order) > 0 {
		query = query.Order(order)
	}

	err := query.Offset(offset).Limit(rows).Find(res).Offset(-1).Limit(-1).Count(count).Error
	if gorm.IsRecordNotFoundError(err) {
		log.Info("record is not found")
		return nil
	} else if err != nil {
		log.Errorf("get failed(%v)", err)
	}

	return err
}
