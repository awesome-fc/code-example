import logging
import os
from mns.account import Account
from mns.topic import *

logger = logging.getLogger()

# initialize
def initialize(context):
    global my_topic
    # 从参数上下文中拿到一组临时密钥，避免了您把自己的AccessKey信息编码在函数代码里
    access_key_id = context.credentials.access_key_id
    access_key_secret = context.credentials.access_key_secret
    security_token = context.credentials.security_token
    # 获取自己 mns 的推送地址 Endpoint
    mns_endpoint = os.getenv("MNS_ENDPOINT")
    # 获取自己 mns topic 的名称
    topic_name = os.getenv("MNS_TOPIC_NAME")
    # 创建mns实例
    my_account = Account(mns_endpoint, access_key_id,
                         access_key_secret, security_token)
    # 获取mns实例的一个Topic对象
    my_topic = my_account.get_topic(topic_name)


def handler(event, context):
    try:
        # 发布一条消息
        msg_body = "I am a test message."
        # 消息正文
        msg = TopicMessage(msg_body)
        # 发送消息
        re_msg = my_topic.publish_message(msg)
        return "Publish Message Succeed. MessageBody:%s MessageID:%s" % (msg_body, re_msg.message_id)
    except MNSExceptionBase as e:
        if e.type == "TopicNotExist":
            logger.info("Topic '{}' not exist, please create topic before send message.".format(topic_name))
        raise RuntimeError(e)