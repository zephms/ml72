package main

import "io/ioutil"

func Checkproject(path string) bool {
	files,_ := ioutil.ReadDir(path)
	for _,i:= range files{
		if i.Name()=="project.xml"{
			return true
		}
	}
	return false
}
