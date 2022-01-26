package intro

import (
	"dev/server"
	"dev/tool"
	"fmt"
	"net"
	"strconv"
)

func serve4client(con *net.Conn, lserver *server.MLserver){
	data := make([]byte, 1000)
	(*con).Read(data[0:2]) // 读取前两位
	nodeLenStr := string(data[0:2])
	nodeLen, _ := strconv.Atoi(nodeLenStr)
	(*con).Read(data[0:nodeLen]) // 读取cmd指令
	fmt.Println("开发到一半 我懒得写了")
}

func receieve(con net.Conn, lserver *server.MLserver){
	data := make([]byte, 1000)
	// 验证头部结构
	currectMlHead := []byte{'M', 'L', '7', '2'}
	_,err := con.Read(data[0:4])
	if err != nil{
		lserver.Logger.DebugInput <- "got head err"
		return
	}
	if !tool.BytesEquals(data[0:4], currectMlHead){
		lserver.Logger.DebugInput <- "got an error head, stoped the connection"
		return
	}
	// 回馈报文
	_, err = con.Write([]byte("MLOK"))
	if err != nil {
		lserver.Logger.DebugInput <- "write error"
		return
	}
	// 添加到客户端列表
	lserver.AddClient(&con)

	// 等待接收指令进程
	go serve4client(&con, lserver)

	//for{
	//	n, err := con.Read(data[0:3])
	//	if err != nil{
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(string(data[0:n]))
	//	time.Sleep(1*time.Second)
	//}
	lserver.Logger.DebugInput <- "one conn closed"
}

func LongServerTestMain(){
	ser := server.NewMLserver()
	address := ser.Ip + ":" + strconv.Itoa(ser.PortNetSocket)
	listen, err := net.Listen("tcp", address)
	if err != nil{
		ser.Logger.DebugInput <- "failed to start tcp server :-("
		fmt.Println(err)
		return
	}
	for {
		con, err := listen.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		ser.Logger.DebugInput <- "got a new conn"
		go receieve(con, ser)
	}
}

