package main

import (
	"fmt"
	"github.com/beevik/etree"
	"strings"
	"sync"
)

func PartLs(){
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./project.xml");err != nil {
		panic(err)
	}
	root := doc.SelectElement("project")
	fmt.Println("here are all run parts by each line:")
	runItems := root.FindElements("./reserve[0]/part")
	for _,i := range runItems{
		fmt.Println(i.FindElement("./name").Text())
	}
}

func RunPart(name string){
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./project.xml");err != nil {
		panic(err)
	}
	root := doc.SelectElement("project")
	parts := root.FindElements("./reserve/part")
	for _,i := range parts {
		nameAttr := i.SelectAttr("name")
		var name string
		if nameAttr != nil {
			name = nameAttr.Value
		} else {
			name = "aPart"
		}
		script := i.FindElement("./script")
		if script != nil {
			str2run := strings.Join(ClearXml(script.Text()), "\n")
			var wg sync.WaitGroup // 其实这里wg并不起作用，为了简略就这么对付了。
			RunCMD(str2run, name, &wg)
			wg.Wait()
		}
	}

}
