## Synopsis
```
import (
	"github.com/ldeng7/gin-utils/logger_duo"
	"github.com/ldeng7/go-x/logx"
)

func main() {
	logger := logger_duo.Init("logs/")
	logger.LogLevel = logx.NOTICE

	logger.Warn("what", 123)
}
