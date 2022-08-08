const { MNSClient } = require('./client');

var client, queue
exports.initialize = (context, callback) => {
    // 设置HTTP协议MNS客户端接入点
    const endpoint = process.env.MnsEndpoint;
    // AccessKey ID阿里云身份验证，在阿里云RAM控制台创建。
    const accessKeyId = context.credentials.accessKeyId;
    // AccessKey Secret阿里云身份验证，在阿里云RAM控制台创建。
    const accessKeySecret = context.credentials.accessKeySecret;
    const queuename = process.env.QueueName;
    client = new MNSClient(endpoint, accessKeyId, accessKeySecret);
    queue = client.getMNSQueue(queuename);
    callback(null, "initialize");
};
exports.handler = async (event, context, callback) => {
    //发送一条消息
    res = await queue.sendMessage("hello mns", "20"); // 20是设置消息20s后可被消费。取值范围：0~604800，单位为秒。 默认为0
    console.log("Send message succ: MessageID:%s,BodyMD5:%s", res.body.MessageId, res.body.MessageBodyMD5);
    callback(null, "succ");
};
