from NetWork import NetWork
import time

'''
v0.1.1
'''


class MLManager:
    def __init__(self, hp) -> None:
        self.host = '127.0.0.1'
        self.port = hp
        pass

    def check(self) -> bool:  # 暂时建议仅限测试使用
        resp = NetWork.pureGet([self.port, "check"])
        if resp == "okok":
            return True
        else:
            return False

    def push(self, channel, d) -> str:
        log = NetWork.loopPost([self.port, "push", channel], {'data': d}, loopFor=NetWork.NETERRorTO)
        return log

    def get(self, channel) -> str:
        resp = NetWork.loopGet([self.port, "get", channel], loopFor=NetWork.NETERRorTO)
        return resp[4:]

    def waitForServer(self) -> None:
        while not self.check():
            print("服务器暂时未启动")
            time.sleep(0.1)
        pass

    def waitForSignal(self, sig, value) -> None:
        passReturn = "okok" + str(value)
        get = NetWork.loopGet([self.port, "getSignal", sig], loopFor=NetWork.ONLY4NETERR)
        while get != passReturn:
            time.sleep(0.05)
            get = NetWork.loopGet([self.port, "getSignal", sig], loopFor=NetWork.ONLY4NETERR)
            print("得到一个不行的", get)

    def setSignal(self, sig, value) -> str:
        log = NetWork.loopPost([self.port, "setSignal", sig], {'data': value}, loopFor=NetWork.ONLY4NETERR)
        return log

    # 用于发送get请求,而不是获得元素,内部方法

    ''' 发送请求之后,没有回应
    可能情况: 服务器未启动 -> 打印提示服务器疑似未启动,并且等待
            服务器未响应 -> 等待响应
            '''


