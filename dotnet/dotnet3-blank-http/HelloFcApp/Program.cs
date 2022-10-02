using System.Threading.Tasks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Aliyun.Serverless.Core;
using Aliyun.Serverless.Core.Http;

namespace Example
{
    public class HttpHandler : FcHttpEntrypoint
    {
        public override async Task<HttpResponse> HandleRequest(HttpRequest request, HttpResponse response,
            IFcContext fcContext)
        {
            response.StatusCode = 200;
            response.ContentType = "text/plain";
            await response.WriteAsync("Hello World\n");
            return response;
        }

        protected override void Init(IWebHostBuilder builder) { }

        static void Main(string[] args) { }
    }
}