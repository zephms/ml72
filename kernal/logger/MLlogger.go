package logger

import (
	"dev/global"
	"fmt"
	"os"
	"strconv"
	"time"
)

var FILENAMEPREFIX = "templog"


type MLlogger struct {
	logPath string
	fp *os.File
	CustomInput chan string
	DebugInput chan string
}

func NewMLlogger(logRoot string) *MLlogger {
	//todo 核验 logroot 是文件夹名
	var mylogger *MLlogger
	if global.Env.Createlog{
		now := time.Now()
		fileName := FILENAMEPREFIX + strconv.FormatInt(now.UnixNano(), 10) + ".mllog.txt"
		filePathAndName := logRoot + "/" + fileName
		//gotfp, err := os.OpenFile(filePathAndName, os.O_WRONLY, 0666)
		gotfp, err := os.Create(filePathAndName) // todo 如果文件存在呢
		if err != nil {
			fmt.Println(err.Error())
			panic("日志系统启动错误~")
		}
		fmt.Println("Log System is Ready...")
		gotfp.WriteString("=== log =======================\n")
		gotfp.WriteString("启动时间："+now.Format("2006-01-02 15:04:05")+"\n")
		mylogger = &MLlogger{
			logPath: filePathAndName,
			fp: gotfp,
			CustomInput: make(chan string),
			DebugInput: make(chan string),
		}
	} else {
		mylogger = &MLlogger{
			logPath: "",
			fp: nil,
			CustomInput: nil,
			DebugInput: nil,
		}
	}

	mylogger.Working()
	return mylogger
}

func (log MLlogger)info(){

}

func (log MLlogger)Working(){
	go func() {
		for{
			select {
			case context := <- log.CustomInput:
				log.fp.WriteString("[MLlogger]\t"+context+"\n")
				fmt.Println("[MLlogger]\t"+context)
			case context := <- log.DebugInput:
				log.fp.WriteString("[debugging]\t"+context+"\n")
				fmt.Println("[debugging]\t"+context)
			}
		}
	}()
}
