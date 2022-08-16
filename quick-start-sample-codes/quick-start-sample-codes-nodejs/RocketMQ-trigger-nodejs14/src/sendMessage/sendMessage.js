const {
  MQClient,
  MessageProperties
} = require('@aliyunmq/mq-http-sdk');

// 设置HTTP协议客户端接入点，进入消息队列RocketMQ版控制台实例详情页面的接入点区域查看。
const endpoint = "${HTTP_ENDPOINT}";
// AccessKey ID阿里云身份验证，在阿里云RAM控制台创建。
const accessKeyId = "${ACCESS_KEY}";
// AccessKey Secret阿里云身份验证，在阿里云RAM控制台创建。
const accessKeySecret = "${SECRET_KEY}";

var client = new MQClient(endpoint, accessKeyId, accessKeySecret);

// 消息所属的Topic，在消息队列RocketMQ版控制台创建。
const topic = "${TOPIC}";
// Topic所属的实例ID，在消息队列RocketMQ版控制台创建。
// 若实例有命名空间，则实例ID必须传入；若实例无命名空间，则实例ID传入null空值或字符串空值。实例的命名空间可以在消息队列RocketMQ版控制台的实例详情页面查看。
const instanceId = "${INSTANCE_ID}";

const producer = client.getProducer(instanceId, topic);

(async function(){
  try {
    // 循环发送4条消息。
    for(var i = 0; i < 4; i++) {
      let res;
      msgProps = new MessageProperties();
      // 设置消息的自定义属性。
      msgProps.putProperty("a", i);
      // 设置消息的Key。
      msgProps.messageKey("MessageKey");
      res = await producer.publishMessage("hello mq.", "", msgProps);
      console.log("Publish message: MessageID:%s,BodyMD5:%s", res.body.MessageId, res.body.MessageBodyMD5);
    }

  } catch(e) {
    // 消息发送失败，需要进行重试处理，可重新发送这条消息或持久化这条数据进行补偿处理。
    console.log(e)
  }
})();