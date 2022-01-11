package run

import (
	"dev/tool"
	"fmt"
	"github.com/beevik/etree"
	"strconv"
	"strings"
	"sync"
)

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
					str2run := strings.Join(tool.ClearXmlinRun(j.Text()), "\r\n")
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