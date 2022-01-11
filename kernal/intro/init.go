package intro

import (
	"fmt"
	"os"
	"dev/tool"
)
import pm "path"

func Initproj(path string)  {
	if tool.Checkproject(path){
		fmt.Println("WARNING! current is a multilang directory already! Do you want to cover it?")
		fmt.Println("Y for cover, N and others for quit. :-)")
		var got string
		fmt.Scan(&got)
		if got=="Y"{
			fmt.Println("welcome to a new project !")
			InitXml(path)
			return
		} else {
			fmt.Println("bye!")
			return
		}
	} else {
		fmt.Println("welcome to a new project !")
		InitXml(path)
		fmt.Println("ok")
	}
}

func InitXml(path string){
	// 这里有一些异常得处理，文件夹是否存在之类的，另外，路径存不存在也是个事
	docpath := pm.Join(path, "project.xml")
	f, _ := os.Create(docpath)
	defer f.Close()
	_, _ = f.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<project>\n    <info>\n        <projName>st</projName>\n        <description>这是一段描述</description>\n    </info>\n    <server>\n        <version>0.1.1</version>\n        <port>8083</port>\n    </server>\n    <runs>\n        <run>\n            <index>debug</index>\n            <part label=\"pyprog\">\n                python testA.py\n            </part>\n            <part label=\"mulline\">\n                python testB.py\n                javac a.java\n                java a\n            </part>\n            <part name=\"py\"/>\n            <part>\n                go run *.go\n            </part>\n        </run>\n        <run>\n            <index>release</index>\n            <part>\n                javac a.java\n                java a\n            </part>\n        </run>\n    </runs>\n    <reserve>\n        <part name=\"py\" label=\"pypart\">\n            <script>python testA.py &amp;&amp; pi&amp;&amp; ng \\\\\\www.baidu.com  pause</script>\n        </part>\n        <part name=\"py\" label=\"xzvc\">\n            <script>python testA.py &amp;&amp; pi&amp;&amp; ng \\\\\\www.baidu.com  pause</script>\n        </part>\n        <part name=\"py\" label=\"fasdsfd\">\n            <script>python testA.py &amp;&amp; pi&amp;&amp; ng \\\\\\www.baidu.com  pause</script>\n        </part>\n    </reserve>\n</project>"))
}