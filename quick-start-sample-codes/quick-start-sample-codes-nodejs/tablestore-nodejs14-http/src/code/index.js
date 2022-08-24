const{ Long, Client } = require("tablestore")

var client;

exports.initialize = (context, callback) => {
  client = new Client({
    accessKeyId: context.credentials.accessKeyId,
    accessKeySecret: context.credentials.accessKeySecret,
    endpoint: process.env.ENDPOINT,
    instancename: process.env.INSTANCE_NAME,
    securityToken: context.credentials.securityToken
  });
  callback(null, "succ")
};

exports.handler = (request, response, context) => {
  // 本示例中表格存储表名为 fc_test, 主键包含两列 region 和 id
  var queries = request.queries
  const tableName = process.env.TABLE_NAME;
  var params = {
    tableName: tableName,
    primaryKey: [{"region": queries['region']}, {"id": Long.fromNumber(queries['id'])}],
    maxVersions: 1,
  };

  response.setHeader('content-type', 'application/json');

  client.getRow(params, (err, res) => {
    if (err) {
      response.setStatusCode(500);
      response.send(Buffer.from(err.message.toString()));
      throw err;
    }
    response.setStatusCode(200);
    response.send(Buffer.from(JSON.stringify(res.row)));
  });
}