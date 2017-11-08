package frame

import (
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	Status  int
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func Render(c *gin.Context, status int, code int, message string, data interface{}) {
	c.JSON(status, &Resp{code, message, data})
}

func RenderDataOrError(c *gin.Context, data interface{}, err *Error) {
	if nil == err {
		c.JSON(200, &Resp{0, "success", data})
	} else {
		c.JSON(err.Status, &Resp{err.Code, err.Message, nil})
	}
}
