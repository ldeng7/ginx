package logger_duo

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/ldeng7/go-x/logx"
)

func Init(logPath string) (*logx.Logger, error) {
	info, err := os.Stat(logPath)
	if nil != err {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(logPath, 0777); nil != err {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else if !info.IsDir() {
		return nil, errors.New("log path is not a dir")
	}

	l, err := logx.Init(&logx.InitArgs{
		Filename: path.Join(logPath, gin.Mode()+".error.log"),
		Flags:    log.LstdFlags | log.Lshortfile,
		LogLevel: logx.INFO,
	})
	if nil != err {
		return nil, err
	}

	la, err := logx.Init(&logx.InitArgs{
		Filename: path.Join(logPath, gin.Mode()+".access.log"),
	})
	if nil != err {
		return nil, err
	}
	gin.DefaultWriter = la.GetWriter()
	return l, nil
}
