# v0.1.1
print("这里是b1")
import mlp as mlp


cm = mlp.MLManager("8083")
cm.waitForServer()

cm.waitForSignal("start", "1")

while True:
    get = cm.get("a-b1")
    print("b1获取到数据:", get)

