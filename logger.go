package ginx

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/go-logger-lite/logger"
)

func NewLogger(writer io.Writer, logPath string, logLevel int) (*logger.Logger, error) {
	l, err := logger.New(&logger.InitArgs{
		Writer:   writer,
		Filename: logPath + ".error.log",
		Flags:    log.LstdFlags | log.Lshortfile,
		LogLevel: logLevel,
	})
	if nil != err {
		return nil, err
	}

	la, err := logger.New(&logger.InitArgs{
		Writer:   writer,
		Filename: logPath + ".access.log",
	})
	if nil != err {
		return nil, err
	}
	gin.DefaultWriter = la.GetWriter()
	return l, nil
}
