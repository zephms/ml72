package intro

import (
	"dev/server"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"strconv"
)


type MLmsg struct {
	cmdStr,path,clientName string
	body string
}

//var defaultCap = 3
//var channelMap map[string]chan string
//var signalMap map[string]string



//func (receiver *MLserver) setSig(ch string) string {
//
//}


// todo 针对ip和端口的绑定是有问题的，没来得及修改
func RunServerN(port int) {

	app := iris.New()
	ser := server.NewMLserver()

	//channelMap = make(map[string]chan string)
	//signalMap := make(map[string]string)

	// 新建一个msg channel
	//var MsgChannel chan MLmsg
	//MsgChannel = make(chan MLmsg)

	// 针对短链接模式，我们选择立即回复or队列回复
	// check
	app.Handle("GET", "/check", func(ctx context.Context) {
		ctx.HTML("ok")
	})

	// channel
	app.Handle("POST", "/push/{ch:string}", func(ctx context.Context) {
		ch := ctx.Params().Get("ch")
		data := ctx.PostValue("data")
		ser.Push(ch, data)
	})

	app.Handle("GET", "/get/{ch:string}", func(ctx context.Context) {
		ch := ctx.Params().Get("ch")
		ctx.HTML("ok"+ ser.Get(ch))
	})

	// signal
	//app.Handle("POST", "/setSignal/{sig:string}", func(ctx context.Context) {
	//	sig := ctx.Params().Get("sig")
	//	data := ctx.PostValue("data")
	//	signalMap[sig] = data
	//	ctx.HTML("ok")
	//})
	//
	//app.Handle("GET", "/getSignal/{sig:string}", func(ctx context.Context) {
	//	sig := ctx.Params().Get("sig")
	//	ctx.HTML("ok"+ signalMap[sig])
	//})

	app.Run(iris.Addr(":"+strconv.Itoa(port)))
}
