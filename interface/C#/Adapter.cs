using System;
using System.Collections.Generic;
using System.IO;
using System.Net;
using System.Text;
using System.Threading;

namespace WinFormsApp1
{
    public class Adapter
    {
        public static int ONLYNETERR = 0;
        public static int NETERRorT0 = 1;

        public static String pureGet(String[] seg)
        {
            String result = "";
            String str2get = "http://127.0.0.1:" + string.Join('/', seg);
            HttpWebRequest req = (HttpWebRequest)WebRequest.Create(str2get);
            HttpWebResponse resp = (HttpWebResponse)req.GetResponse();
            Stream stream = resp.GetResponseStream();
            try
            {
                using (StreamReader reader = new StreamReader(stream))
                {
                    result = reader.ReadToEnd();
                    return "ok" + result;
                }
            }
            finally 
            { 
                stream.Close();
            }
           
        }

        public static String purePost(String[] seg, Dictionary<string, string> dic)
        {
            String str2post = "http://127.0.0.1:" + string.Join('/', seg);

            string result = "";
            HttpWebRequest req = (HttpWebRequest)WebRequest.Create(str2post);
            req.Method = "POST";
            req.ContentType = "application/x-www-form-urlencoded";
            #region 添加Post 参数
            StringBuilder builder = new StringBuilder();
            int i = 0;
            foreach (var item in dic)
            {
                if (i > 0)
                    builder.Append("&");
                builder.AppendFormat("{0}={1}", item.Key, item.Value);
                i++;
            }
            byte[] data = Encoding.UTF8.GetBytes(builder.ToString());
            req.ContentLength = data.Length;
            using (Stream reqStream = req.GetRequestStream())
            {
                reqStream.Write(data, 0, data.Length);
                reqStream.Close();
            }
            #endregion
            HttpWebResponse resp = (HttpWebResponse)req.GetResponse();
            Stream stream = resp.GetResponseStream();
            //获取响应内容
            using (StreamReader reader = new StreamReader(stream, Encoding.UTF8))
            {
                result = reader.ReadToEnd();
            }
            return "ok" + result;

        }
    }
}


