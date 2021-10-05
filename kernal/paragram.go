package main

import (
	"fmt"
	"strconv"
)

func AnaParagram(ags []string)  {
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
		if len(ags)==2 {
			Initproj(".")
		} else {
			if IsExistDir(ags[2]){
				Initproj(ags[2])
			} else {
				fmt.Println("the directory does not exist, please check again. :-(")
			}
		}
	} else if ags[1] == "run" {
		if len(ags)==2 {
			if Checkproject("."){
				RunLs()
			} else {
				fmt.Println("the current directory may not a multilang project, change or init.")
			}
		} else if len(ags) == 3{
			if Checkproject("."){
				RunRun(ags[2])
			} else {
				fmt.Println("the current directory may not a multilang project, change or init.")
			}
		} else {
			fmt.Println("unexpected arguments, please read doc :-) ")
		}
	} else if ags[1] == "server" {
		if len(ags)==2{
			// runserver random port
			port := GetPort(8080)
			RunServer(port)
		} else if len(ags)==3 {
			// run server ags[3
			if ags[2] == "project" {

			} else {
				port, err := strconv.Atoi(ags[2])
				if err != nil {
					fmt.Println("the port would be numerical")
					return
				} else {
					RunServer(port)
				}
			}
		}
	} else if ags[1] == "build" {
		fmt.Println("功能还在开发中。。。")
	} else if ags[1] == "part" {
		fmt.Println("功能还在开发中，，，")
	}
}
