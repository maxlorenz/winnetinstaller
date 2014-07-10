using System;
using System.Net;

namespace Consumer
{
    class MyWebClient : WebClient
    {
        protected override WebRequest GetWebRequest(Uri address)
        {
            WebRequest w = base.GetWebRequest(address);
            w.Timeout = 20 * 60 * 1000;
            w.Proxy = null;

            return w;
        }

        public int TimeOut { set; get; }
    }
}
