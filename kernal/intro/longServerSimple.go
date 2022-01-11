package intro

import(
	"dev/server"
	"fmt"
    "net"
	"strconv"
	"time"
)

func Read(con net.Conn, lserver *server.MLserver){
	data := make([]byte, 1000)
	for{
		n, err := con.Read(data[0:3])
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(string(data[0:n]))
		time.Sleep(1*time.Second)
	}
	lserver.Logger.DebugInput <- "one conn closed"
}

func LongServerTestMain(){
	ser := server.NewMLserver()
	address := ser.Ip + ":" + strconv.Itoa(ser.Port)
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
		go Read(con, ser)
	}
}

