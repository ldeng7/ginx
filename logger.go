package ginx

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/go-logx/logx"
)

func InitLogger(writer io.Writer, logPath string, logLevel int) (*logx.Logger, error) {
	l, err := logx.Init(&logx.InitArgs{
		Writer:   writer,
		Filename: logPath + ".error.log",
		Flags:    log.LstdFlags | log.Lshortfile,
		LogLevel: logLevel,
	})
	if nil != err {
		return nil, err
	}

	la, err := logx.Init(&logx.InitArgs{
		Writer:   writer,
		Filename: logPath + ".access.log",
	})
	if nil != err {
		return nil, err
	}
	gin.DefaultWriter = la.GetWriter()
	return l, nil
}
