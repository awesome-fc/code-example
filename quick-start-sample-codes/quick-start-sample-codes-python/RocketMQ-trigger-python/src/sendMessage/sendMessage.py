import sys

from mq_http_sdk.mq_exception import MQExceptionBase
from mq_http_sdk.mq_producer import *
from mq_http_sdk.mq_client import *
import time

# 初始化client。
mq_client = MQClient(
    # 设置HTTP协议客户端接入点，进入消息队列RocketMQ版控制台实例详情页面的接入点区域查看。
     "${HTTP_ENDPOINT}",
    # AccessKey ID阿里云身份验证，在阿里云RAM控制台创建。
     "${ACCESS_KEY}",
    # AccessKey Secret阿里云身份验证，在阿里云RAM控制台创建。
    "${SECRET_KEY}"
    )
# 消息所属的Topic，在消息队列RocketMQ版控制台创建。
topic_name = "${TOPIC}"
# Topic所属的实例ID，在消息队列RocketMQ版控制台创建。
# 若实例有命名空间，则实例ID必须传入；若实例无命名空间，则实例ID传入空字符串。实例的命名空间可以在消息队列RocketMQ版控制台的实例详情页面查看。
instance_id = "${INSTANCE_ID}"

producer = mq_client.get_producer(instance_id, topic_name)

# 循环发送2条消息。
msg_count = 2
print("%sPublish Message To %s\nTopicName:%s\nMessageCount:%s\n" % (10 * "=", 10 * "=", topic_name, msg_count))

try:
    for i in range(msg_count):
            msg = TopicMessage(
                    # 消息内容。
                    "I am test message %s.hello" % i,
                    # 消息标签。
                    "tag %s" % i
                        )
            # 设置消息的自定义属性。
            msg.put_property("a", i)
            # 设置消息的Key。
            msg.set_message_key("MessageKey")
            re_msg = producer.publish_message(msg)
            print("Publish Message Succeed. MessageID:%s, BodyMD5:%s" % (re_msg.message_id, re_msg.message_body_md5))

except MQExceptionBase as e:
    if e.type == "TopicNotExist":
         print("Topic not exist, please create it.")
         sys.exit(1)
    print("Publish Message Fail. Exception:%s" % e)