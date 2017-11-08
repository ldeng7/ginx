package gorm_util

import (
	"github.com/jinzhu/gorm"
)

type GormCtx struct {
	Db *gorm.DB
}

func (g *GormCtx) New() *gorm.DB {
	return g.Db.New()
}

func (g *GormCtx) First(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	out := db.First(value, where...)
	if (nil != out.Error) && !out.RecordNotFound() {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Last(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	out := db.Last(value, where...)
	if (nil != out.Error) && !out.RecordNotFound() {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Find(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	out := db.Find(value, where...)
	if (nil != out.Error) && !out.RecordNotFound() {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Count(db *gorm.DB, value interface{}) *gorm.DB {
	out := db.Count(value)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Scan(db *gorm.DB, dest interface{}) *gorm.DB {
	out := db.Scan(dest)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Create(db *gorm.DB, value interface{}) *gorm.DB {
	out := db.Create(value)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Save(db *gorm.DB, value interface{}) *gorm.DB {
	out := db.Save(value)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Update(db *gorm.DB, attrs ...interface{}) *gorm.DB {
	out := db.Update(attrs...)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Updates(db *gorm.DB, values interface{}, ignore ...bool) *gorm.DB {
	out := db.Updates(values, ignore...)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) UpdateColumn(db *gorm.DB, attrs ...interface{}) *gorm.DB {
	out := db.UpdateColumn(attrs...)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) UpdateColumns(db *gorm.DB, values interface{}) *gorm.DB {
	out := db.UpdateColumns(values)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}

func (g *GormCtx) Delete(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	out := db.Delete(value, where...)
	if nil != out.Error {
		panic(out.Error)
	}
	return out
}
