package server

import (
	"dev/logger"
)

type MLserver struct {
	Db     map[string]MLnode
	Logger logger.MLlogger

	Ip   string
	Port int
	//l log.Logger
}

func NewMLserver() *MLserver {
	nullDb := new(map[string]MLnode)
	//newlogger := Logger.NewMLlogger("D:\\sssss\\mllog")
	newlogger := logger.NewMLlogger(".")
	return &MLserver{
		Db:     *nullDb,
		Logger: *newlogger,

		Ip:   "127.0.0.1",
		Port: 8080,
	}
}

func (receiver *MLserver) Push(ch string, data string)  {
	node, ok := receiver.Db[ch]

	if ok {
		if node.MLtype == MLNodeType_CHANNEL {
			curChan := node.Data.(chan string)
			curChan <- data
		} else {

			//receiver.Db[ch] <- data
			receiver.Logger.DebugInput <- "try for this channel, but it is used"
		}
	} else {
		var newChan = make(chan string, MLNodeChanDefaultCap)
		var newChanNode = NewMLnode(ch, MLNodeType_CHANNEL, newChan)
		receiver.Db[ch] = *newChanNode
		receiver.Logger.DebugInput <- "create new node, type is channel"
	}
}



func (receiver *MLserver) Get(ch string) string {
	var res string
	node, ok := receiver.Db[ch]

	if ok {
		if node.MLtype == MLNodeType_CHANNEL {
			curChan := node.Data.(chan string)
			res = <- curChan
		} else {

			//receiver.Db[ch] <- data
			receiver.Logger.DebugInput <- "try for this channel, but it is used"
			panic("get了类型不匹配的")
		}
	} else {
		var newChan = make(chan string, MLNodeChanDefaultCap)
		var newChanNode = NewMLnode(ch, MLNodeType_CHANNEL, newChan)
		receiver.Db[ch] = *newChanNode
		res = <- receiver.Db[ch].Data.(chan string)
		receiver.Logger.DebugInput <- "create new node, type is channel"
	}
	return res

}