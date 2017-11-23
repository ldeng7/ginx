package gorm_util

import (
	"github.com/jinzhu/gorm"
)

type GormError struct {
	error
}

type GormCtx struct {
	Db *gorm.DB
}

func (g *GormCtx) New() *gorm.DB {
	return g.Db.New()
}

// CRUDs

func outOrPanic(out *gorm.DB) *gorm.DB {
	if nil != out.Error {
		panic(GormError{out.Error})
	}
	return out
}

func outOrNotfoundOrPanic(out *gorm.DB) *gorm.DB {
	if (nil != out.Error) && !out.RecordNotFound() {
		panic(GormError{out.Error})
	}
	return out
}

func (g *GormCtx) First(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	return outOrNotfoundOrPanic(db.First(value, where...))
}

func (g *GormCtx) Last(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	return outOrNotfoundOrPanic(db.Last(value, where...))
}

func (g *GormCtx) Find(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	return outOrNotfoundOrPanic(db.Find(value, where...))
}

func (g *GormCtx) Count(db *gorm.DB, value interface{}) *gorm.DB {
	return outOrPanic(db.Count(value))
}

func (g *GormCtx) Scan(db *gorm.DB, dest interface{}) *gorm.DB {
	return outOrPanic(db.Scan(dest))
}

func (g *GormCtx) Create(db *gorm.DB, value interface{}) *gorm.DB {
	return outOrPanic(db.Create(value))
}

func (g *GormCtx) Save(db *gorm.DB, value interface{}) *gorm.DB {
	return outOrPanic(db.Save(value))
}

func (g *GormCtx) Update(db *gorm.DB, attrs ...interface{}) *gorm.DB {
	return outOrPanic(db.Update(attrs...))
}

func (g *GormCtx) Updates(db *gorm.DB, values interface{}, ignore ...bool) *gorm.DB {
	return outOrPanic(db.Updates(values, ignore...))
}

func (g *GormCtx) UpdateColumn(db *gorm.DB, attrs ...interface{}) *gorm.DB {
	return outOrPanic(db.UpdateColumn(attrs...))
}

func (g *GormCtx) UpdateColumns(db *gorm.DB, values interface{}) *gorm.DB {
	return outOrPanic(db.UpdateColumns(values))
}

func (g *GormCtx) Delete(db *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	return outOrPanic(db.Delete(value, where...))
}
