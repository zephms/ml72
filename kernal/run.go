package main

import (
	"fmt"
	"github.com/beevik/etree"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func RunLs(){
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./project.xml");err != nil {
		panic(err)
	}
	root := doc.SelectElement("project")
	fmt.Println("here are all run configurations by each line:")
	runItems := root.FindElements("./runs[0]/run")
	for _,i := range runItems{
		fmt.Println(i.FindElement("./index").Text())
	}
}

func RunRun(index string){
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./project.xml");err != nil {
		panic(err)
	}
	root := doc.SelectElement("project")
	runItems := root.FindElements("./runs[0]/run")
	var wg sync.WaitGroup
	for _,i := range runItems{ // 对每一个run进行遍历， i是一个run节点
		if i.FindElement("./index").Text() == index {
			defaultCount := 1
			for _, j := range i.FindElements("./part"){ // 对i进行遍历 j是part节点
				name := j.SelectAttr("name")
				if name != nil {
					fmt.Println("执行part", name.Value)
					fmt.Println("这里应该执行这个模块，但是由于我来不及了，下个版本再写；要是用到的话，不要使用name引用方式了，直接放里面就好")
				} else {
					label := j.SelectAttr("label")
					str2run := strings.Join(ClearXml(j.Text()), "\r\n")
					if label != nil {
						// label run
						wg.Add(1)
						go RunCMD(str2run, label.Value, &wg)
					} else {
						// defalt run
						wg.Add(1)
						go RunCMD(str2run, "default"+strconv.Itoa(defaultCount), &wg)
					}
				}
			}
			break
		}
	}
	wg.Wait()
}

// ClearXml 将一个xml内容，清洗成一行一行的，这里可能会去掉一行最开始和最后的空格; 空号交给run cmd去解决吧
func ClearXml(s string) []string {
	ss := strings.Split(s, "\n")
	for i:=0;i<len(ss);i++ {
		reg := regexp.MustCompile("^[ \u3000\t\r\n]*")
		ss[i] = reg.ReplaceAllString(ss[i], "")

		reg = regexp.MustCompile("[ \u3000\t\r\n]*$")
		ss[i] = reg.ReplaceAllString(ss[i], "")
	}
	return ss
}