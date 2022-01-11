package intro

import (
	"dev/run"
	"dev/tool"
	"fmt"
	"strconv"
)

func AnaParagram(ags []string)  {
	if len(ags) ==1 {
		tool.ShowLogo()
		fmt.Println("usage：ml <mode> [args...]")
		fmt.Println("")
		fmt.Println("where <mode> is one of:")
		fmt.Println("    init, run, part, server, help")
		fmt.Println("")
		fmt.Println("\tinit    ml init")
		fmt.Println("\t        ml init [<path>]")
		fmt.Println("")
		fmt.Println("\trun     ml run [<project name>]")
		fmt.Println("")
		fmt.Println("\tpart    ml part [<run module name>]")
		fmt.Println("")
		fmt.Println("\tserver  ml server")
		fmt.Println("\t        ml server [<port>]")
		fmt.Println("")
		fmt.Println("for more information, please visit: https://ml.acbs.top")
		return
	}
	if ags[1]=="init"{
		if len(ags)==2 {
			Initproj(".")
		} else {
			if tool.IsExistDir(ags[2]){
				Initproj(ags[2])
			} else {
				fmt.Println("the directory does not exist, please check again. :-(")
			}
		}
	} else if ags[1] == "run" {
		if len(ags)==2 {
			if tool.Checkproject("."){
				run.RunLs()
			} else {
				fmt.Println("the current directory may not a multilang project, change or init.")
			}
		} else if len(ags) == 3{
			if tool.Checkproject("."){
				run.RunRun(ags[2])
			} else {
				fmt.Println("the current directory may not a multilang project, change or init.")
			}
		} else {
			fmt.Println("unexpected arguments, please read doc :-) ")
		}
	} else if ags[1] == "server" {
		if len(ags)==2{
			// runserver random port
			port := tool.GetPort(8080)
			RunServerN(port)
		} else if len(ags)==3 {
			// run server ags[3
			if ags[2] == "project" {

			} else {
				port, err := strconv.Atoi(ags[2])
				if err != nil {
					fmt.Println("the port would be numerical")
					return
				} else {
					RunServerN(port)
				}
			}
		}
	} else if ags[1] == "build" {
		fmt.Println("功能还在开发中。。。")
	} else if ags[1] == "part" {
		fmt.Println("功能还在开发中，，，")
	} else if ags[1] == "lserver"{
		fmt.Println("启动long server。。。")
		fmt.Println("功能还在开发中，，，")
	}
}