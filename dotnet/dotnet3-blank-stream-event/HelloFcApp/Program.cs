using System.IO;
using System.Threading.Tasks;
using Aliyun.Serverless.Core;
using Microsoft.Extensions.Logging;

namespace Example
{
    public class Hello
    {
        public async Task<Stream> StreamHandler(Stream input, IFcContext context)
        {
            ILogger logger = context.Logger;
            logger.LogInformation("Handle request: {0}", context.RequestId);
            MemoryStream copy = new MemoryStream();
            await input.CopyToAsync(copy);
            copy.Seek(0, SeekOrigin.Begin);
            return copy;
        }

        static void Main(string[] args){}
    }
}