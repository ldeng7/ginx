package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RES_CODE_OK            = 0
	RES_CODE_GENERAL_ERROR = 1
)

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type RespError struct {
	Status  int
	Code    int
	Message string
}

func (e *RespError) Error() string {
	return e.Message
}

type Context struct {
	*gin.Context
}

func (c *Context) RenderData(data interface{}) {
	c.JSON(http.StatusOK, &Resp{RES_CODE_OK, "success", data})
}

func (c *Context) RenderError(err error) {
	if respErr, ok := err.(*RespError); ok {
		if nil != respErr {
			status, code := respErr.Status, respErr.Code
			if 0 == status {
				status = http.StatusOK
			}
			if 0 == code {
				code = RES_CODE_GENERAL_ERROR
			}
			c.JSON(status, &Resp{code, respErr.Message, nil})
		} else {
			c.JSON(http.StatusOK, &Resp{RES_CODE_GENERAL_ERROR, "error", nil})
		}
	} else if nil != err {
		c.JSON(http.StatusOK, &Resp{RES_CODE_GENERAL_ERROR, err.Error(), nil})
	} else {
		c.JSON(http.StatusOK, &Resp{RES_CODE_GENERAL_ERROR, "error", nil})
	}
}

func (c *Context) RenderDataOrError(data interface{}, err error) {
	if nil == err {
		c.RenderData(data)
	} else {
		c.RenderError(err)
	}
}

type ICreateController interface {
	Create(*Context)
}

type IListController interface {
	List(*Context)
}

type IGetController interface {
	Get(*Context)
}

type IUpdateController interface {
	Update(*Context)
}

type IDeleteController interface {
	Delete(*Context)
}

func AddRoutes(group *gin.RouterGroup, path string, controller interface{}) {
	if con, ok := controller.(ICreateController); ok {
		group.POST(path, func(gc *gin.Context) {
			con.Create(&Context{gc})
		})
	}
	if con, ok := controller.(IListController); ok {
		group.GET(path, func(gc *gin.Context) {
			con.List(&Context{gc})
		})
	}
	if con, ok := controller.(IGetController); ok {
		group.GET(path+"/:id", func(gc *gin.Context) {
			con.Get(&Context{gc})
		})
	}
	if con, ok := controller.(IUpdateController); ok {
		group.PUT(path+"/:id", func(gc *gin.Context) {
			con.Update(&Context{gc})
		})
	}
	if con, ok := controller.(IDeleteController); ok {
		group.DELETE(path+"/:id", func(gc *gin.Context) {
			con.Delete(&Context{gc})
		})
	}
}
