using Aliyun.Serverless.Core;
using Microsoft.Extensions.Logging;

namespace Example
{
    public class Hello
    {
        public class Product
        {
            public string Id { get; set; }

            public string Description { get; set; }
        }

        // optional serializer class, if it’s not specified, the default serializer (based on JSON.Net) will be used.
        // [FcSerializer(typeof(MySerialization))]
        public Product PocoHandler(Product product, IFcContext context)
        {
            string Id = product.Id;
            string Description = product.Description;
            context
                .Logger
                .LogInformation("Id {0}, Description {1}", Id, Description);
            return product;
        }

        static void Main(string[] args)
        {
        }
    }
}
