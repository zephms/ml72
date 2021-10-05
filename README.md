<p class="tip">(很遗憾,之前答应帮我优化文档的hxd最近一直很忙，因此，这个文档虽然还只是草稿版，还是不得不顶上来了)</p>

# 写在前面

我很久之前在各种项目中，每个人会的技术不同，而且我们使用不同的技术完成不同的东西，想往一起捏合是一个非常费劲的事情，比如我们使用nodejs-electron做的界面，同时我们还有python-pytorch做的识别核心，两者之间交互的问题我们想了很久，无论怎么整都难以配置难以打包发布，最后我们使用了一种很难看的方式实现了，打包过程也变得非常变态。还有一些其他的项目经历，渐渐地使我清晰了需求，我想要一个东西，可以让我方便地用多种技术实现一个项目，快捷省事地打包发布，甚至希望可以很方便地在演示电脑上快速配置好我们后台依赖已经乱成一锅粥的项目。这个东西的每个模块或许很多框架实现了很多次，包括各种消息队列，RPC框架，微服务系列，redis... 每一个都是倾注了无数多的心血打磨出来的，而我在技术上只能算班门弄斧，项目也很单薄。不过，"既然踩过的坑没法避开，那就把它踩实了。" 我希望用我的方式，让编程者简单容易地实现快速部署和打包，不要为了一点小事还学习几天大框架。


# 简介
> 本项目是一个可以方便地在一个项目中使用多种编程语言、多个进程的代码库. 该平台赋予了小型团队使用不同编程语言分别开发,快速成型电脑端程序的能力.
本项目包括:
- 启动器: 通过配置文件 <kbd>project.xml</kbd> 快速启动不同进程的模块

- 服务程序: 在本地搭建一个HTTP服务器,用于数据的处理和交流

- 客户端接口程序: 每种编程语言都会配一个接口类,用于实现傻瓜式数据交互

- 测试器: 多进程交互项目独立开发时,需模拟其他程序的传入参数,因此提供了测试器,可以简单地编写自动化的参数传输方式。换句话说，为了方便开发的工具


# 原理简述
> 我们在本地搭建一个服务进程,利用服务进程完成数据交流
数据的交流我们目前定义了两种形式,分别是 <kbd>channel</kbd> 、 <kbd>signal</kbd>，另外兼容socket模式。

- channel: 我们模拟了一个一个管道,一个进程向某个管道输入信息,而另一个进程从另一个进程取信息

- signal: 如同名字一样,这是一个信号,一个进程设置了一个信号,其他的进程就可以读取到信号.与channel不用的是,signal可以一次设置而多次访问.

至于为什么选用这两种呢？我当初是这么想的：在网络上，人和人的交流，只要通过人对人消息，群聊，打电话，公告这四种方式就可以了。可能从数学原理上还有缺陷，但是日常能想到的就这么几种，也算先从方便用户理解开始叭，后续有需要再加。所以 发消息,群聊->channel, 打电话->socket, 公告->signal. (当然这里我知道人和人之间的沟通是一种模糊的沟通，需要附带大量的确认环节，而程序间沟通是精确的，不必附带大量验证而浪费时间空间，所以这么模拟意义似乎不是很大，但是我也思考了一些实际的例子，实际用到的和想到的场景使用这个沟通逻辑确实丝滑流畅)

# 快速开始

推荐下载ml.exe作为项目管理工具，并添加到环境变量
```cmd
wget 未来的下载链接
```
成功后，在cmd中执行命令 ml 会得到一个简略的文档

```cmd
PS D:\sssss\multilanguage\docs> ml

============================================

              ____  _   __
  __ _  __ __/ / /_(_) / /  ___ ____  ___ _
 /  ' \/ // / / __/ / / /__/ _ `/ _ \/ _ `/
/_/_/_/\_,_/_/\__/_/ /____/\_,_/_//_/\_, / 
                                    /___/  

        multiLang   @  ml.acbs.top
         Welcome to use MultiLang
============================================

usage：ml <mode> [args...]

where <mode> is one of:
    init, run, part, server, help

        init    ml init
                ml init [<path>]

        run     ml run [<project name>]

        part    ml part [<run module name>]

        server  ml server
                ml server [<port>]

for more information, please visit: https://ml.acbs.top
```

## 初始化项目

通过<kbd>init</kbd>初始化项目。初始化成功后，可以看到在当前目录下创建了该项目入口文件 <kbd>project.xml</kbd>

project.xml中配置了不同的程序入口，每一个入口有包括好几个进程

每一个进程的启动方式用一段cmd命令表示。

## 测试运行

使用管理工具通过<kbd>run</kbd>模式进行运行，不带其他参数表示以运行表第一种运行方案运行
```cmd
ml run
```
按照其他的预定方案运行：
```cmd
ml run debug
```
其中 <kbd>debug</kbd> 为 <kbd>project.xml</kbd> 中的 <kbd>index</kbd> 块内容


## 配置project.xml

使用管理工具初始化得到的project.xml是这样的

```xml
<?xml version="1.0" encoding="utf-8"?>
<project>
    <info>
        <projName>st</projName>
        <description>这是一段描述</description>
    </info>
    <server>
        <version>0.1.1</version>
        <port>8083</port>
    </server>
    <runs>
        <run>
            <index>debug</index>
            <part label="pyprog">
                python testA.py
            </part>
            <part label="mulline">
                python testB.py
                javac a.java
                java a
            </part>
            <part name="py"/>
            <part>
                go run *.go
            </part>
        </run>
        <run>
            <index>release</index>
            <part>
                javac a.java
                java a
            </part>
        </run>
    </runs>
    <reserve>
        <part name="py" label="pypart">
            <script>python testA.py &amp;&amp; pi&amp;&amp; ng \\\www.baidu.com  pause</script>
        </part>
        <part name="py1" label="xzvc">
            <script>python testA.py &amp;&amp; pi&amp;&amp; ng \\\www.baidu.com  pause</script>
        </part>
        <part name="py2" label="fasdsfd">
            <script>python testA.py &amp;&amp; pi&amp;&amp; ng \\\www.baidu.com  pause</script>
        </part>
    </reserve>
</project>
```

该xml原则来讲使用utf8进行编码。根节点project表示该目录，project下的server表示服务器端的版本号，和本项目使用的端口号。project下的runs表示本项目的一些入口。project下的reserve为预设的一些进程。其中的part就是一个进程。通过 ml part name即可启动某一进程。如下：

```cmd
ml part py
```

让我们仔细看一下runs这些入口。
每一个入口为一个run。run块以一个index名为开始，其余为一些part。index为启动入口的索引，而run代表了一个进程。每一个run块在启动的时候，会根据名字查看是否在reserve中有储存的，如果有的话，就会执行reserve中的part块。label属性用于启动起来之后，标记输出来自于哪一个进程。

## 启动项目

执行如下cmd命令：
```cmd
ml run default
```

## 发布

使用管理工具，<kbd>build</kbd> 得到一个bat文件，以后运行该文件即可
例如：
```cmd
ml build release
```
即可得到一个文件：projectName_release.bat，运行该脚本即可完成项目的启动

# 接口用法

根据接口

## HTTP Api

 - 以后补上

## Python

 - 以后补上

## GO，java, C#, matlab...

 - 以后补上


# 例子

## Step 1 : 启动服务器

```cmd
ml serve 8083

```


> 第二步和第三步我们模拟了一个主程序两个从程序的单向发送数据的情况,主程序产生随机数, 将奇数通过a-b1通道传递给从程序b1.py,将偶数通过a-b2通道传递给从程序b2.go

## Step 2 : 端口发送数据:
``` python - a.py
import random
import time
import interface.mlp as mlp

cm = mlp.MLManager('8083')
cm.waitForServer()

get = cm.setSignal("start", "1")


while True:
    ri = random.randint(0, 1000)
    print("这次得到的值是", ri)
    if ri % 2 == 1:
        get = cm.push("a-b1", str(ri))
    if ri % 2 == 0:
        get = cm.push("a-b2", str(ri))
    time.sleep(1)
```

## Step 3 : 接收数据:
``` python - b1.py
print("这里是b1")
import interface.mlp as mlp


cm = mlp.MLManager("8083")
cm.waitForServer()

cm.waitForSignal("start", "1")

while True:
    get = cm.get("a-b1")
    print("b1获取到数据:", get)
```

``` go - b2.go
package main

import (
	"./mlg"
	"fmt"
)

func main () {
	cm := mlg.NewMLManager("8083")
	cm.WaitForServer()

	cm.WaitForSignal("start", "1")

	for {
		get := cm.Get("a-b2")
		fmt.Println("b2获取到的数据", get)
	}

}
```


未完待续...