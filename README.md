
# 写在前面

在大量的项目开发中，很多开发者会遇到一个问题————项目应该用什么语言开发？不同语言有不同语言的好处，有的语言更贴近底层，但却没有提供足够严肃的开发规范和统一框架；有的语言对网络层并发做得更好，但在客户端交互部分却已经被淘汰了些许年头，也许你甚至很难在最近四年的博客中找到相关技术？有的语言几乎直接隔绝了开发者和底层，但却提供了极多的框架和接口，对于快速开发来说非常友好。<br>
决定我们的开发语言，一般是开发者需要做的第一件，也许会决定整个项目走向，甚至导致后续开发路程坎坷or顺利的决策，并非选择简单、接口模块足够多的胶水语言开发就会简单，很大程度上项目开发语言的选择需要领导者有足够的远视，能一开始就大致决定好开发路径。而很多时候，我们也不得不使多种语言混合使用，比如最简单的——C/Cpp编写dll文件，再用python调用，或许是因为你需要使用真多线程，或许是你对python某部分的性能要求更高......总而言之，很多时候，多语言协作，是我们不得不面临的一块儿小骨头————不难，却极麻烦。<br>

# 目前的多语言协同模式

笔者在目前的开发中，认为多语言协同工作模式可以大致分这三种：<br>
<br>
1、打包编译后文件，例如前文提到的dll文件，so文件等，使得此语言的一些特性可以随时被调用，好处就是简单，目前大多数软件也确实这样做的，但坏处也很明显————开发者不太可能把已经做好的整体运行模块拆分出功能部分和运行部分，再把功能部分给打包成dll文件，对于小型团体/个人开发者来说代价太大。<br>
<br>
2、进程间通信，开发者完全可以让某语言进行开始进程，然后唤醒各个语言的进程，再使他们用socket在本地回环端口进行通信，虽然不在一个解释器中执行，但却可以通过通信来达到步骤匹配的目的，此方法的好处更明显————一个优秀的开发者，完全可以把进程玩得很得心应手，多进程协同工作也不算什么难事。但缺点也是有的，毕竟不是每个人都有很高的水平，可以随便把进程和线程玩弄股掌之间，就算可以，每次针对性编写进程通信逻辑，是不是也很烧脑呢？<br>
<br>
3、微服务通信，此方法和进程间通信大同小异，区别就在于不要求本地通信，即便不同进程在不同服务器上也可以进行通信。<br>
<br>
本项目就针对于多语言协作通信问题，开发出了本框架，试图提供一个简单的开发规范接口，使大多开发者可以快速略过这一部分。<br>

# 简介
> 本项目是一个可以方便地在一个项目中使用多种编程语言、多个进程的代码库. 该平台赋予了小型团队使用不同编程语言分别开发,快速成型电脑端程序的能力.
本项目包括:
- 启动器: 通过配置文件 <kbd>project.xml</kbd> 快速启动不同进程的模块

- 服务程序: 在本地搭建一个HTTP服务器,用于数据的处理和交流

- 客户端接口程序: 每种编程语言都会配一个接口类,用于实现傻瓜式数据交互

- 测试器: 多进程交互项目独立开发时,需模拟其他程序的传入参数,因此提供了测试器,可以简单地编写自动化的参数传输方式。换句话说，为了方便开发的工具

- 兼容器 \ 虚拟机 (新版本开发中，功能尚未明确，后续再定)


# 原理简述
我们在本地搭建一个服务进程,利用服务进程完成数据交流。数据的交流我们目前定义了四种形式,分别是 <kbd>channel</kbd> 、 <kbd>signal</kbd>、<kbd>dataBind</kbd>、<kbd>rpc</kbd> ， 兼容五种数据交换方式：HTTP请求、socket、系统管道、文本文件(万不得已舍弃性能时再用)、内存共享(下一步再实现)

- channel: 我们模拟了一个一个管道,一个进程向某个管道输入信息,而另一个进程从另一个进程取信息

- signal: 如同名字一样,这是一个信号,一个进程设置了一个信号,其他的进程就可以读取到信号.与channel不用的是,signal可以一次设置而多次访问.

- dataBind：两个程序同时操作同一个变量，当变量被一个其他线程改变时，自动触发钩子函数。目前仅兼容字符串，int32，字节串三种。

- rpc：跟众多rpc框架差不多，舍弃了一部分功能

> 兼容多种数据交流方式，是因为使得不同的进程可以根据自己编程语言语法的方便程度随意挑选，比如nodejs模块对服务程序发送http请求，而python模块使用socket与服务程序连接。注意，实际性能上是有差别的，等我研究研究性能测试，会大致推荐一下。

<p class="tip">后两种在实现中，改完bug再传到仓库里</p>


# 快速开始

推荐下载ml.exe作为项目管理工具，并添加到环境变量
```cmd
wget http://未来的下载链接/ml.exe
```
成功后，在cmd中执行命令 ml 会得到一个简略的文档

```cmd
PS D:\sssss\ml\docs> ml

============================================

              ____  _   __
  __ _  __ __/ / /_(_) / /  ___ ____  ___ _
 /  ' \/ // / / __/ / / /__/ _ `/ _ \/ _ `/
/_/_/_/\_,_/_/\__/_/ /____/\_,_/_//_/\_, / 
                                    /___/  

            ML   @  ml.acbs.top
              Welcome to use
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

使用管理工具初始化得到的project.xml。经过修改是这样的

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

让我们仔细看一下runs这些入口。<br>
每一个入口为一个run。<br>
run块以一个index名为开始，其余为一些part。<br>
index为启动入口的索引，而run代表了一个进程。每一个run块在启动的时候，会根据名字查看是否在reserve中有储存的，如果有的话，就会执行reserve中的part块。<br>
label属性用于启动起来之后，标记输出来自于哪一个进程。<br>

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

以后补上

## HTTP Api

 - 以后补上

## Python

 - 以后补上

## GO，java, C#, matlab...

 - 以后补上


# 例子

## 试验一下

### Step 1 : 启动服务器

```cmd
ml serve 8083

```


> 第二步和第三步我们模拟了一个主程序两个从程序的单向发送数据的情况,主程序产生随机数, 将奇数通过a-b1通道传递给从程序b1.py,将偶数通过a-b2通道传递给从程序b2.go

### Step 2 : 端口发送数据:
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

### Step 3 : 接收数据:
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

## 插入大佬的代码

（介绍rpc）一段python程序，中间有一个功能难以实现，多方请求下，一个大佬提供了matlab的这个功能代码，而这个模块要执行很多次，难以手动解决

#### 代码在整理中。。。

## electron-python融合程序

（介绍数据绑定） 没想好

## 高效爬虫

（综合）golang具有比python更好的网络io性能，因此使用golang来完成大规模io访问，python分析网络中获得的json，由C#完成界面开发

#### 代码在整理中。。。

# 鸣谢

衷心感谢莫皓对本项目文档撰写的大力支持！

未完待续...

