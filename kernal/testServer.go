package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"strconv"
)

func runTServer(port int){
	app := iris.New()

	app.Handle("GET", "/check", func(ctx context.Context) {
		ctx.HTML("ok")
	})

	// 变量绑定：
	//app.Handle("")

	app.Run(iris.Addr(":"+strconv.Itoa(port)))
}
