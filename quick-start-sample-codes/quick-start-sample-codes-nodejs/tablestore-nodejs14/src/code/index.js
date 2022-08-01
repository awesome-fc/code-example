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

exports.handler = (event, context, callback) => {
  // 本示例中表格存储表名为 fc_test, 主键包含两列 region 和 id
  const tableName = process.env.TABLE_NAME;
  var params = {
    tableName: tableName,
    primaryKey: [{"region": "abc"}, {"id": Long.fromNumber(1)}],
    maxVersions: 1,
  };

  client.getRow(params, (err, res) => {
    if (err) {
      callback(err); 
      throw err;
    }
    callback(null, res.row)
  });
}