package intro

import (
	logger2 "dev/logger"
	"fmt"
)

func Maindevtest(){
	fmt.Println("hhhhhhh")
	logger := logger2.NewMLlogger(".")
	logger.CustomInput <- "for cust"
	logger.DebugInput <- "hello world"
	//for  {
	//	time.Sleep(2*time.Second)
	//}
}
