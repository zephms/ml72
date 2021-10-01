# 主程序
# v0.1.1
import random
import time
import mlp as mlp

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


