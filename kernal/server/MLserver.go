package server

import (
	"dev/logger"
	"dev/server/base"
	"dev/server/nodes"
	"dev/tool"
	"net"
)

type MLserver struct {
	Db     map[string]server.MLnode
	Logger logger.MLlogger

	Ip            string
	PortNetSocket int
	PortIris int

	longClients map[string]nodes.MLclient
}

func NewMLserver() *MLserver {
	nullDb := new(map[string]server.MLnode)
	//newlogger := Logger.NewMLlogger("D:\\sssss\\mllog")
	newlogger := logger.NewMLlogger(".")
	return &MLserver{
		Db:     *nullDb,
		Logger: *newlogger,

		Ip:            "127.0.0.1",
		PortNetSocket: 8081,
		PortIris: 8080,
	}
}

func (receiver *MLserver) AddClient(con *net.Conn)  {
	id := tool.GetDiffTimeStampStr()
	newClient := nodes.NewMLclient(id, con, receiver)
	receiver.longClients[id] = *newClient
}

func (receiver MLserver) Exec(nodePath string, cmd []byte) []byte {
	node, ok := receiver.Db[nodePath]
	if ok {
		return node.Exec(cmd)
	} else {
		return []byte("no such node")
	}
}

func (receiver *MLserver) Push(ch string, data string)  {
	node, ok := receiver.Db[ch]

	if ok {
		if node.MLtype == nodes.MLNodeType_CHANNEL {
			curChan := node.Data.(chan string)
			curChan <- data
		} else {

			//receiver.Db[ch] <- data
			receiver.Logger.DebugInput <- "try for this channel, but it is used"
		}
	} else {
		var newChan = make(chan string, nodes.MLNodeChanDefaultCap)
		var newChanNode = server.NewMLnode(ch, nodes.MLNodeType_CHANNEL, newChan)
		receiver.Db[ch] = *newChanNode
		receiver.Logger.DebugInput <- "create new node, type is channel"
	}
}

func (receiver *MLserver) Get(ch string) string {
	var res string
	node, ok := receiver.Db[ch]

	if ok {
		if node.MLtype == nodes.MLNodeType_CHANNEL {
			curChan := node.Data.(chan string)
			res = <- curChan
		} else {

			//receiver.Db[ch] <- data
			receiver.Logger.DebugInput <- "try for this channel, but it is used"
			panic("get了类型不匹配的")
		}
	} else {
		var newChan = make(chan string, nodes.MLNodeChanDefaultCap)
		var newChanNode = server.NewMLnode(ch, nodes.MLNodeType_CHANNEL, newChan)
		receiver.Db[ch] = *newChanNode
		res = <- receiver.Db[ch].Data.(chan string)
		receiver.Logger.DebugInput <- "create new node, type is channel"
	}
	return res

}
