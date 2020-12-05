package ginx

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/go-x/logx"
)

func NewLogger(writer io.Writer, logPath string, logLevel int) (*logx.Logger, error) {
	l, err := logx.New(&logx.InitArgs{
		Writer:   writer,
		Filename: logPath + ".error.log",
		Flags:    log.LstdFlags | log.Lshortfile,
		LogLevel: logLevel,
	})
	if nil != err {
		return nil, err
	}

	la, err := logx.New(&logx.InitArgs{
		Writer:   writer,
		Filename: logPath + ".access.log",
	})
	if nil != err {
		return nil, err
	}
	gin.DefaultWriter = la.GetWriter()
	return l, nil
}
