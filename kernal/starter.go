package main

import (
	"fmt"
	"os"

)
func showLogo(){
		fmt.Println("============================================\n")
		ascii := "              ____  _   __                 \n  __ _  __ __/ / /_(_) / /  ___ ____  ___ _\n /  ' \\/ // / / __/ / / /__/ _ `/ _ \\/ _ `/\n/_/_/_/\\_,_/_/\\__/_/ /____/\\_,_/_//_/\\_, / \n                                    /___/  \n"
		fmt.Println(ascii)
		fmt.Println("        multiLang   @  ml.acbs.top")
		fmt.Println("         Welcome to use MultiLang")
		fmt.Println("============================================\n")
}


func main1()  {
	ags := os.Args
	if len(ags) ==1 {
		showLogo()
		fmt.Println("usage：ml <mode> [args...]")
		fmt.Println("")
		fmt.Println("where <mode> is one of:")
		fmt.Println("    init, run, part, server, help")
		fmt.Println("")
		fmt.Println("\tinit\tml init")
		fmt.Println("\t\t\tml init [<path>]")
		fmt.Println("")
		fmt.Println("\trun\t\tml run [<project name>]")
		fmt.Println("")
		fmt.Println("\tpart\tml part [<run module name>]")
		fmt.Println("")
		fmt.Println("\tserver\tml server")
		fmt.Println("\t\t\tml server [<port>]")
		fmt.Println("")
		fmt.Println("for more information, please visit: https://ml.acbs.top")
		return
	}
	if ags[1]=="init"{
		if len(ags)==3 {
			Initproj(".")
		} else {
			// init ags[3
			// todo 判断是否是路径的功能交给了initproj函数
			Initproj(ags[3])
		}
	} else if ags[1] == "run" {
		if len(ags)==4 {
			// run mode ags[3
		} else {
			fmt.Println("unexpected arguments, please read doc :-) ")
		}
	} else if ags[1] == "server" {
		if len(ags)==3{
			// runserver random port
		} else if len(ags)==4 {
			// run server ags[3
		}
	} else if ags[1] == "build" {
		fmt.Println("功能还在开发中。。。")
	}
}

func main()  {
	Initproj("./test")
}