# 从程序2
# v0.1.1
print("这里是b2")
import mlp as mlp



cm = mlp.MLManager('8083')
cm.waitForServer()

cm.waitForSignal("start", "1")

while True:
    get = cm.get("a-b2")
    print("b2获取到数据:", get)