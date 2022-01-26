package intro

import (
	"dev/server"
	"fmt"
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
func RunServerN(port int, ser *server.MLserver) {

	app := iris.New()
	//ser := server.NewMLserver()

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

	// simple


	simpleRouter := app.Party("/simple")

	simpleRouter.Get("/{name:string}/home", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("you name: %s",name)
	})
	// route: /user/post
	simpleRouter.Post("/post", func(ctx iris.Context) {
		ctx.Writef("method:%s,path;%s",ctx.Method(),ctx.Path())
	})

	app.Get("/download", func(ctx context.Context) {
		src := "D:/a/aa.txt"
		ctx.SendFile(src, "code")
	})

	app.Post("/ml/exec", func(ctx iris.Context) {
		bodyBytes, err := ctx.GetBody()
		if err!= nil {
			return
		}
		fmt.Println(bodyBytes)
		ctx.Write(ser.Exec("/ml/exec", bodyBytes))
		//nodePath := ctx.Params().Get("nodePath")
		//ctx.WriteString(ser.Exec(nodePath, cmd))
	})


	app.RegisterView(iris.HTML("./pages", ".html"))
	app.Get("/ml/webcmd", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		//ctx.ViewData("message", "Hello world!")
		// 渲染模板文件： ./views/hello.html
		ctx.View("cmd.html")
	})
	app.Get("/", func(ctx context.Context) {
		ctx.View("navigate.html")
	})

	app.Run(iris.Addr(":"+strconv.Itoa(port)),iris.WithCharset("UTF-8"))

}
