import logging
import sys
import os
from mns.account import Account
from mns.queue import *

logger = logging.getLogger()


def initialize(context):
    global my_queue
    # 从参数上下文中拿到一组临时密钥，避免了您把自己的AccessKey信息编码在函数代码里
    access_key_id = context.credentials.access_key_id
    access_key_secret = context.credentials.access_key_secret
    security_token = context.credentials.security_token
    # 获取自己 mns 的推送地址 Endpoint
    mns_endpoint = os.getenv("MnsEndpoint")
    # 获取自己 mns queue 的名称
    queue_name = os.getenv("QueueName")
    # 创建mns实例
    my_account = Account(mns_endpoint, access_key_id,
                         access_key_secret, security_token)
    # 获取mns实例的一个 queue 对象
    my_queue = my_account.get_queue(queue_name)


def handler(event, context):

    # 发布一条消息
    msg_body = "I am a test message."
    # 消息正文
    msg = Message(msg_body)
    # 发送消息
    re_msg = my_queue.send_message(msg)
    return("Send Message Succeed. MessageBody:%s MessageID:%s" % (msg_body, re_msg.message_id))
