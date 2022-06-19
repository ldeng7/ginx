package main

import (
	"os"
	"strconv"

	"github.com/ldeng7/ginx/examples/server/context"
)

func main() {
	if err := run(); err != nil {
		println(err.Error())
	}
}

func run() error {
	if err := context.Init(os.Args[1]); err != nil {
		return err
	}
	ctx := context.Instance()
	r := setRoutes()
	return r.Run(":" + strconv.Itoa(ctx.Config.HttpPort))
}
