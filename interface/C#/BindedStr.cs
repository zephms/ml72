using System;
using System.Collections.Generic;
using System.Text;
using System.Threading;

namespace WinFormsApp1
{
    public class BindedStr
    {
        MLclient mc;
        String last;
        Action fun;
        Thread thread;

        public BindedStr(String last, Action fun)
        {
            this.last = last;
            this.fun = fun;

            thread = new Thread(new ThreadStart(work));
            //thread.Start();
        }

        public void work()
        {
            
        }

    }
}
