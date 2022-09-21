'use strict';
const {
    MQClient,
    MessageProperties
} = require('@aliyunmq/mq-http-sdk');

var client, producer
exports.initialize = (context, callback) => {
    const accessKeyId = context.credentials.accessKeyId;
    const accessKeySecret = context.credentials.accessKeySecret;
    const sercurityToken = context.credentials.securityToken;
    const endpoint = process.env.ROCKETMQ_ENDPOINT;
    const instanceId = process.env.INSTANCEID;
    const topic = process.env.TOPIC;
    client = new MQClient(endpoint, accessKeyId, accessKeySecret, sercurityToken)
    producer = client.getProducer(instanceId, topic);
    callback(null,"initialize");
}
exports.handler = async (event,context,callback) => {
    try {
        console.log("Start to send message");
        // Parse the json
        let res;
        let msgProps = new MessageProperties();
        msgProps.putProperty("a", 1);
        msgProps.messageKey("MessageKey");
        res = await producer.publishMessage("hello RocketMQ!", "", msgProps);
        console.log("Publish message: MessageID:%s,BodyMD5:%s", res.body.MessageId, res.body.MessageBodyMD5);
        callback(null,"send message success");
    } catch(e){
        console.log(e);
        callback(null,"send message error");
    }
};