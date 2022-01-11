package run

import (
	"fmt"
	"github.com/beevik/etree"
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
