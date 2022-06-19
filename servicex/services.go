package servicex

import (
	"github.com/ldeng7/ginx/ginx"
	"gorm.io/gorm"
)

type ITableNamerModel interface {
	TableName() string
}

type IResponserModel interface {
	Resp() map[string]any
}

func CheckDup(db *gorm.DB, model any, msg string) error {
	if db = db.Last(model); db.Error == gorm.ErrRecordNotFound {
		return nil
	} else if db.Error != nil {
		return db.Error
	}
	return &ginx.RespError{Code: ginx.RES_CODE_ALREADY_EXIST, Message: msg}
}

func Get(db *gorm.DB, id uint, model any, msg string) error {
	if db = db.Last(model, id); db.Error == gorm.ErrRecordNotFound {
		return &ginx.RespError{Code: ginx.RES_CODE_NOT_FOUND, Message: msg}
	} else if err := db.Error; nil != err {
		return err
	}
	return nil
}

func GetEx(db *gorm.DB, model any, msg string) error {
	if db = db.Last(model); db.Error == gorm.ErrRecordNotFound {
		return &ginx.RespError{Code: ginx.RES_CODE_NOT_FOUND, Message: msg}
	} else if err := db.Error; nil != err {
		return err
	}
	return nil
}

type ListResp struct {
	List  []map[string]any `json:"list"`
	Total int64            `json:"total"`
}

func List[T IResponserModel](models *[]T, dbList, dbTotal *gorm.DB) (*ListResp, error) {
	var err error
	if err = dbList.Find(models).Error; nil != err {
		return nil, err
	}

	l := make([]map[string]any, len(*models))
	for i, model := range *models {
		l[i] = model.Resp()
	}

	r := &ListResp{l, 0}
	if nil != dbTotal {
		if err := dbTotal.Count(&r.Total).Error; nil != err {
			return nil, err
		}
	} else {
		r.Total = int64(len(l))
	}
	return r, nil
}
