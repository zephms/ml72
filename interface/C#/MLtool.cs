using System;
using System.Collections.Generic;
using System.Threading;

namespace WinFormsApp1
{
    class MLtool
    {
        String host;
        String port;
        MLclient mlkernal;

        public MLtool(String port1)
        {
            host = "127.0.0.1";
            port = port1;
            mlkernal = new MLclient(host, port);
        }

        public bool check()
        {
            String resp = Adapter.pureGet(new String[] { port, "check" });
            if (resp == "okok")
            {
                return true;
            }
            else
            {
                return false;
            }
        }

        public String push(String channel, String d)
        {
            String log = loopPost(new string[] { port, "push", channel }, new Dictionary<string, string> { { "data", d } }, Adapter.NETERRorT0);
            return log;
        }

        public String get(String channel)
        {
            String resp = loopGet(new string[] { port, "get", channel }, Adapter.NETERRorT0);
            return resp.Substring(4);
        }

        public void waitForServer()
        {
            while (!check())
            {
                Console.WriteLine("服务器暂时未启动");
                Thread.Sleep(100);
            }
        }

        public void waitForSignal(String sig, String value)
        {
            String passReturn = "okok" + value;
            String get = loopGet(new string[] { port, "getSignal", sig }, Adapter.ONLYNETERR);
            while (get != passReturn)
            {
                Thread.Sleep(50);
                Console.WriteLine("得到一个不行的");
                get = loopGet(new string[] { port, "getSignal", sig }, Adapter.ONLYNETERR);

            }
        }

        public String setSignal(String sig, String value)
        {
            String log = loopPost(new string[] { port, "setSignal", sig }, new Dictionary<string, string> { { "data", value } }, Adapter.ONLYNETERR);
            return log;
        }


        /// tools:
        public static String loopGet(String[] seg, int loopFor)
        {
            bool loopFlag = true;
            String get = "er8798";
            while (loopFlag)
            {
                get = Adapter.pureGet(seg);
                if (loopFor == Adapter.ONLYNETERR)
                {
                    if (!get.StartsWith("er"))
                    {
                        loopFlag = false;
                    }
                    else
                    {
                        if (get.StartsWith("er") | get.StartsWith("okto"))
                        {
                            loopFlag = true;
                        }
                        else
                        {
                            loopFlag = false;
                        }
                    }
                    Thread.Sleep(50);
                }
            }
            return get;
        }

        public static String loopPost(String[] seg, Dictionary<string, string> dic, int loopFor)
        {
            bool loopFlag = true;
            String get = "er0987098";
            while (loopFlag)
            {
                get = Adapter.purePost(seg, dic);
                if (loopFor == Adapter.ONLYNETERR)
                {
                    if (!get.StartsWith("er"))
                    {
                        loopFlag = false;
                    }
                    else
                    {
                        if (get.StartsWith("er") | get.StartsWith("okto"))
                        {
                            loopFlag = true;
                        }
                        else
                        {
                            loopFlag = false;
                        }
                    }
                    Thread.Sleep(50);
                }
            }
            return get;
        }


    }

    
}
