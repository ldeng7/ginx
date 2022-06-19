package servicex

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Criteria[T ITableNamerModel] struct {
	Db    *gorm.DB
	conds map[string]any
	model T
}

func NewCriteria[T ITableNamerModel](db *gorm.DB, model T) *Criteria[T] {
	return &Criteria[T]{
		Db:    db,
		conds: map[string]any{},
		model: model,
	}
}

func (c *Criteria[T]) WhereId(key string, id uint) {
	if id != 0 {
		c.conds[key] = id
	}
}

func (c *Criteria[T]) WhereString(key, s string) {
	if len(s) != 0 {
		c.conds[key] = s
	}
}

func (c *Criteria[T]) WhereDate(key string, date *time.Time) {
	if date != nil {
		c.conds[key] = date.Format("2006-01-02")
	}
}

func (c *Criteria[T]) WhereStringFrag(key, s string) {
	if len(s) != 0 {
		c.Db = c.Db.Where(key+" LIKE ?", "%"+s+"%")
	}
}

func (c *Criteria[T]) WhereIds(key string, ids []uint) {
	if len(ids) != 0 {
		c.Db = c.Db.Where(key+" IN (?)", ids)
	}
}

func (c *Criteria[T]) WhereNonIds(key string, ids []uint) {
	if len(ids) != 0 {
		c.Db = c.Db.Where(key+" NOT IN (?)", ids)
	}
}

func (c *Criteria[T]) WhereM2MId(joint T, jointKey, jointAssoKey string, jointAssoId uint) {
	if jointAssoId != 0 {
		jointName := joint.TableName()
		c.Db = c.Db.Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.id AND %s.%s = %d",
			jointName, jointName, jointKey, c.model.TableName(), jointName, jointAssoKey, jointAssoId))
	}
}

func (c *Criteria[T]) Compile(limit, offset int) (*gorm.DB, *gorm.DB) {
	db := c.Db
	var dbt *gorm.DB
	if len(c.conds) != 0 {
		db = db.Where(c.conds)
	}
	if limit != 0 {
		dbt = db.Table(c.model.TableName())
		db = db.Limit(limit).Offset(offset)
	}
	return db, dbt
}
