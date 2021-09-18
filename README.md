# 简介
> 本项目是一个可以方便地在一个项目中使用多种编程语言,多个进程的代码库. 该平台赋予了小型团队使用不同编程语言分别开发,快速成型电脑端程序的能力.
本项目包括:
 - 启动器: 参考了idea系列的 run configuration功能, 通过 配置文件/ui界面 ,一键启动整个进程

- 服务程序: 在本地搭建一个HTTP服务器,用于数据的处理和交流

- 客户端接口程序: 每种编程语言都会配一个接口类,用于实现数据交互

- 测试器: 多进程交互项目独立开发时,需模拟其他程序的传入参数,因此提供了测试器,可以简单地编写自动化的参数传输方式


# 原理简述
> 我们在本地搭建一个服务进程,利用服务进程完成数据交流
数据的交流我们目前定义了两种形式,分别是 <kbd>channel</kbd> 和 <kbd>signal</kbd> 

- channel: 我们模拟了一个一个管道,一个进程向某个管道输入信息,而另一个进程从另一个进程取信息

- signal: 如同名字一样,这是一个信号,一个进程设置了一个信号,其他的进程就可以读取到信号.与channel不用的是,signal可以一次设置而多次访问.


# 快速开始

## Step 1 : 启动服务器

<p class="tip">(很遗憾,这一步还没有做的很人性化,我们依旧需要一个独立运行的服务进程,而且无法友好地将端口号自动地传输给其他的进程,未来一定有机会)</p>

```python 
if __name__ == '__main__':
    port = 8083
    # while check_port_in_use(port=port):
    #     port+=1
    startServer(str(port))
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