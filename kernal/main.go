package main

import (
	"dev/intro"
	"dev/server"
)

func main() {
	//fmt.Println("welcome")
	//tool.ShowLogo()


	//intro.AnaParagram(os.Args)
	//intro.LongServerTestMain()

	ser := server.NewMLserver()
	intro.RunServerN(ser.PortIris, ser)


	//intro.Maindevtest()
}



