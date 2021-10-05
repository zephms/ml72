package main

import (
	"fmt"
	"github.com/beevik/etree"
	"path"
)

func GetXmlDoc(pathstr string) *etree.Document {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path.Join(pathstr, "project.xml"));err != nil {
		panic(err)
	}
	return doc
}

func AnaProjectXml(pathstr string) {
	fmt.Println("语法分析在开发中。。。")

}
