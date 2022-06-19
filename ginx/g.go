package ginx

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	RES_CODE_OK = iota
	RES_CODE_GENERAL_ERROR
	RES_CODE_INVALID_REQ
	RES_CODE_NOT_FOUND
	RES_CODE_ALREADY_EXIST
	RES_CODE_NOT_AUTHED
	RES_CODE_NOT_ALLOWED
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type RespError struct {
	Status  int
	Code    int
	Message string
}

func (e *RespError) Error() string {
	return e.Message
}

type G struct {
	*gin.Context
}

func (g *G) ParsePathId(key string) (uint, error) {
	id64, _ := strconv.ParseUint(g.Param(key), 10, 64)
	if id64 == 0 {
		return 0, &RespError{Code: RES_CODE_INVALID_REQ}
	}
	return uint(id64), nil
}

func (g *G) ParseQueryId(key string) uint {
	id64, _ := strconv.ParseUint(g.Query(key), 10, 64)
	return uint(id64)
}

func (g *G) ParseQueryDate(key string) *time.Time {
	if date, err := time.Parse("2006-01-02", g.Query(key)); nil == err {
		return &date
	}
	return nil
}

func (g *G) ParseQueryIdList(key string) []uint {
	s := g.Query(key)
	if len(s) == 0 {
		return nil
	}
	ids := []uint{}
	idStrs := strings.Split(s, ",")
	for _, idStr := range idStrs {
		id, _ := strconv.ParseUint(idStr, 10, 64)
		if id != 0 {
			ids = append(ids, uint(id))
		}
	}
	return ids
}

func (g *G) ParseQueryPageArgs(maxSize uint64) (limit, offset uint) {
	pageSize, _ := strconv.ParseUint(g.Query("page_size"), 10, 64)
	page, _ := strconv.ParseUint(g.Query("page"), 10, 64)
	if pageSize != 0 && page != 0 {
		if pageSize > maxSize {
			pageSize = maxSize
		}
		limit = uint(pageSize)
		offset = uint(page-1) * limit
	}
	return
}

func (g *G) RenderData(data any) {
	g.JSON(http.StatusOK, &Resp{RES_CODE_OK, "success", data})
}

func (g *G) RenderMessage(code int, message string) {
	g.JSON(http.StatusOK, &Resp{code, message, nil})
}

func (g *G) RenderError(err error) {
	if respErr, ok := err.(*RespError); ok {
		if nil != respErr {
			status, code := respErr.Status, respErr.Code
			if status == 0 {
				status = http.StatusOK
			}
			if code == 0 {
				code = RES_CODE_GENERAL_ERROR
			}
			g.JSON(status, &Resp{code, respErr.Message, nil})
		} else {
			g.JSON(http.StatusOK, &Resp{RES_CODE_OK, "success", nil})
		}
	} else if nil != err {
		g.JSON(http.StatusOK, &Resp{RES_CODE_GENERAL_ERROR, err.Error(), nil})
	} else {
		g.JSON(http.StatusOK, &Resp{RES_CODE_OK, "success", nil})
	}
}

func (g *G) RenderDataOrError(data any, err error) {
	if nil == err {
		g.RenderData(data)
	} else {
		g.RenderError(err)
	}
}
