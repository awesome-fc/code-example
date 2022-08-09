const { MNSClient } = require('./client');

var client, queue
exports.initialize = (context, callback) => {
    // 设置MNS接入点和队列名称
    const endpoint = process.env.MNS_ENDPOINT;
    const queuename = process.env.MNS_QUEUE_NAME;

    // AccessKey ID/AccessKey Secret 阿里云身份验证，在阿里云RAM控制台创建。
    const accessKeyId = context.credentials.accessKeyId;
    const accessKeySecret = context.credentials.accessKeySecret;

    client = new MNSClient(endpoint, accessKeyId, accessKeySecret);
    queue = client.getMNSQueue(queuename);
    callback(null, "initialize succ");
};

exports.handler = async (req, resp, context) => {
    //发送一条消息
    res = await queue.sendMessage("hello mns", "2"); // 2是设置消息2s后可被消费。取值范围：0~604800，单位为秒。 默认为0
    console.log("Send message succ: MessageID:%s,BodyMD5:%s", res.body.MessageId, res.body.MessageBodyMD5);
    resp.setStatusCode(200)
    resp.setHeader("Content-Type", "text/plain");
    resp.send('Send message succ, MessageID:' + res.body.MessageId);
};
