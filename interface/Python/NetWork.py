import requests
import time
# v0.1.1

class NetWork:
    ONLY4NETERR = 0
    NETERRorTO = 1

    # @classmethod
    # def getTool(cls, loop, *seg):
    #     str2get = 'http://127.0.0.1:' + "/".join(seg)
    #     if loop:
    #         while True:
    #             try:
    #                 get = requests.get(str2get)
    #                 return get
    #             except Exception:
    #                 print("我也不知道什么错误")
    #             time.sleep(0.05)
    #     else:
    #         try:
    #             get = requests.get(str2get)
    #             err = None
    #         except ConnectionError:
    #             print('connection error, 暂时怀疑是服务程序未启动')
    #             get = None
    #             err = "err7345843"
    #         return get, err
    #
    # @classmethod
    # def postTool(cls, loop, *seg, **data):
    #     str2post = 'http://127.0.0.1:' + "/".join(seg)
    #     if loop:
    #         while True:
    #             try:
    #                 get = requests.post(str2post, data)
    #                 return get
    #             except Exception:
    #                 print(str2post)
    #                 print("一个错误")
    #                 time.sleep(0.05)
    #                 raise
    #
    #     else:
    #         try:
    #             get = requests.post(str2post, data)
    #             err = None
    #         except ConnectionError:
    #             print("connection err yiwaichucuo")
    #             get = None
    #             err = "err87809"
    #         return get, err

    @classmethod
    def pureGet(cls, seg) -> str:
        str2get = 'http://127.0.0.1:' + "/".join(seg)
        try:
            get = requests.get(str2get)
            return "ok" + get.text
        except ConnectionError:
            return "er87098"
        except Exception:
            return "er9887980"

    @classmethod
    def purePost(cls, seg, data) -> str:
        str2post = 'http://127.0.0.1:' + "/".join(seg)
        try:
            get = requests.post(str2post, data)
            return "ok" + get.text
        except ConnectionError:
            return "er76987"
        except Exception:
            return "er9834"

    @classmethod
    def loopPost(cls, seg, data, loopFor) -> str:
        loopFlag = True
        get = "er5667"
        while loopFlag:
            get = NetWork.purePost(seg, data)
            if loopFor == NetWork.ONLY4NETERR:
                if get[:2] != "er":
                    loopFlag = False
            else:
                if get[:2] == "er" or get[:4] == "okto":
                    loopFlag = True
                else:
                    loopFlag = False
            time.sleep(0.05)
        return get

    @classmethod
    def loopGet(cls, seg, loopFor) -> str:
        loopFlag = True
        get = "er98096"
        while loopFlag:
            get = NetWork.pureGet(seg)
            if loopFor == NetWork.ONLY4NETERR:
                if get[:2] != "er":
                    loopFlag = False
            else:
                if get[:2] == "er" or get[:4] == "okto":
                    loopFlag = True
                else:
                    loopFlag = False
            time.sleep(0.05)
        return get
