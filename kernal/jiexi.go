package main

import (
	"fmt"
	"github.com/beevik/etree"
)

// xml中的转义
// 存储式运行

func runserver(port string){
	fmt.Println("runserver tidai", port)
}

func execCmd(cmd string, showWin bool){
	fmt.Println("exec ", cmd, showWin)
}

func main() {
	fmt.Println("============================================\n")
	ascii := "              ____  _   __                 \n  __ _  __ __/ / /_(_) / /  ___ ____  ___ _\n /  ' \\/ // / / __/ / / /__/ _ `/ _ \\/ _ `/\n/_/_/_/\\_,_/_/\\__/_/ /____/\\_,_/_//_/\\_, / \n                                    /___/  \n"
	fmt.Println(ascii)
	fmt.Println("        multiLang   @  ml.acbs.top")
	fmt.Println("         Welcome to use MultiLang")
	fmt.Println("============================================\n")
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./project.xml");err != nil {
		panic(err)
	}
	root := doc.SelectElement("project")

	projName := root.FindElement("./info[0]/projName").Text()
	fmt.Println("[info] Project Name:", projName)

	desc := root.FindElement("./info[0]/description").Text()
	fmt.Println("[info] Description:", desc)

	serverPort := root.FindElement("./server[0]/port").Text()
	fmt.Println("[info] Running on:", serverPort)

	go runserver(serverPort)

	//f, err := os.Create("test.txt")
	//if err!=nil {
	//	panic(err)
	//}
	runs := root.FindElements("./runs[0]/run")
	for _,run := range runs{
		cd:=run.FindElement("./showcmd[0]")
		fmt.Println(cd)
		//fmt.Println(temp=="")
	}

}