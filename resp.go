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
	ResCode int         `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type RespError struct {
	StatusCode int
	ResCode    int
	Message    string
}

func (e *RespError) Error() string {
	return e.Message
}

func RenderData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Resp{RES_CODE_OK, "success", data})
}

func RenderError(c *gin.Context, err error) {
	if respErr, ok := err.(*RespError); ok {
		if nil != respErr {
			statusCode, resCode := respErr.StatusCode, respErr.ResCode
			if 0 == statusCode {
				statusCode = http.StatusOK
			}
			if 0 == resCode {
				resCode = RES_CODE_GENERAL_ERROR
			}
			c.JSON(statusCode, &Resp{resCode, respErr.Message, nil})
		} else {
			c.JSON(http.StatusOK, &Resp{RES_CODE_GENERAL_ERROR, "error", nil})
		}
	} else if nil != err {
		c.JSON(http.StatusOK, &Resp{RES_CODE_GENERAL_ERROR, err.Error(), nil})
	} else {
		c.JSON(http.StatusOK, &Resp{RES_CODE_GENERAL_ERROR, "error", nil})
	}
}

func RenderDataOrError(c *gin.Context, data interface{}, err error) {
	if nil == err {
		RenderData(c, data)
	} else {
		RenderError(c, err)
	}
}
