package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"strconv"
)

var defaultCap = 3
var channelMap map[string]chan string
var signalMap map[string]string



func push(ch string, data string)  {
	curChan, ok := channelMap[ch]
	if ok{
		curChan <- data
	} else {
		channelMap[ch] = make(chan string, defaultCap)
		//curChan = make(chan string, defaultCap)
		channelMap[ch] <- data
	}
}

func get(ch string) string {
	curChan, ok := channelMap[ch]
	var res string
	if ok{
		res = <- curChan
	} else {
		channelMap[ch] = make(chan string, defaultCap)
		res = <- channelMap[ch]
	}
	return res
}

func RunServer(port int) {

	app := iris.New()
	channelMap = make(map[string]chan string)
	signalMap = make(map[string]string)
	// check
	app.Handle("GET", "/check", func(ctx context.Context) {
		ctx.HTML("ok")
	})
	// channel
	app.Handle("POST", "/push/{ch:string}", func(ctx context.Context) {
		ch := ctx.Params().Get("ch")
		data := ctx.PostValue("data")
		push(ch, data)
	})

	app.Handle("GET", "/get/{ch:string}", func(ctx context.Context) {
		ch := ctx.Params().Get("ch")
		ctx.HTML("ok"+get(ch))
	})

	// signal
	app.Handle("POST", "/setSignal/{sig:string}", func(ctx context.Context) {
		sig := ctx.Params().Get("sig")
		data := ctx.PostValue("data")
		signalMap[sig] = data
		ctx.HTML("ok")
	})

	app.Handle("GET", "/getSignal/{sig:string}", func(ctx context.Context) {
		sig := ctx.Params().Get("sig")
		ctx.HTML("ok"+signalMap[sig])
	})


	//port := getPort(8080)
	//err := os.Setenv("ML", strconv.Itoa(port))
	//if err != nil {
	//	fmt.Println("程序注册失败")
	//	return
	//}
	app.Run(iris.Addr(":"+strconv.Itoa(port)))
}

