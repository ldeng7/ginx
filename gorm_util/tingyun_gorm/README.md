## Synopsis

```
import (
	tingyun "github.com/TingYunAPM/go"
	"github.com/sudiyi/sdy/tingyun_sdy"
	"github.com/ldeng7/gin-utils/gorm_util/tingyun_gorm"
)

var g *tingyun_gorm.Gorm

func init() {
	// you should create a tingyun_gorm.Gorm instance in a global context
	// NOT in each time requesting
	g, err = tingyun_gorm.NewGorm("user:password@tcp(127.0.0.1:3306)/db")
}

// Your service
func ServeA(action *tingyun.Action) {
	db := g.NewDb(action, "your operation name")
	// Now you get a *gorm.DB, use it as usual!
	db.Model(user).UpdateColumn("mobile", "123")
}
```
