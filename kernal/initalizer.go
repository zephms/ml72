package main

import (
	"fmt"
	"os"
)
import pm "path"

func writeProjectXml(path string){
	// 这里有一些异常得处理，文件夹是否存在之类的，另外，路径存不存在也是个事
	docpath := pm.Join(path, "project.xml")
	f, _ := os.Create(docpath)
	defer f.Close()
	_, _ = f.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<project>\n    <info>\n        <projName>st</projName>\n        <description>这是一段描述</description>\n    </info>\n    <server>\n        <version>0.1.1</version>\n        <port>8083</port>\n    </server>\n    <runs>\n        <run>\n            <index>debug</index>\n            <debug>true</debug>\n            <part>\n                <showcmd>true</showcmd>\n                <cd/>\n                <script>python testA.py</script>\n            </part>\n        </run>\n        <run>\n            <index>release</index>\n            <debug>false</debug>\n            <part>\n                <showcmd>false</showcmd>\n                <cd/>\n                <script>python testA.py</script>\n            </part>\n        </run>\n    </runs>\n    <reserve>\n        <part>\n            <name>py</name>\n            <cd ></cd>\n            <script>python testA.py &amp;&amp; ping www.baidu.com && pause</script>\n        </part>\n    </reserve>\n</project>\n"))
}

func Initproj(path string)  {
	if Checkproject(path){
		fmt.Println("WARNING! current is a multilang directory already! Do you want to cover it?")
		fmt.Println("Y for cover, N and others for quit. :-)")
		var got string
		fmt.Scan(&got)
		if got=="Y"{
			fmt.Println("welcome to a new project !")
			writeProjectXml(path)
			return
		} else {
			fmt.Println("bye!")
			return
		}
	} else {
		fmt.Println("hhh")
	}
}


