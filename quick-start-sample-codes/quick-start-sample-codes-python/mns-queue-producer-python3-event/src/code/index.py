"""
本代码样例主要实现以下功能:
* 从环境变量中解析出 MNS 队列模型的配置信息
* 根据以上获取的信息，初始化 MNS 客户端
* 向 MNS 写入一条测试数据


This sample code is mainly doing the following things:
* Get MNS configuration information from environment variables
* Initiate MNS client
* Send a test message to MNS Queue

"""
# -*- coding: utf-8 -*-
import logging
import os
from mns.account import Account
from mns.queue import *

my_queue = None
queue_name = ""

def initialize(context):
    global my_queue, queue_name
    # 从参数上下文中拿到一组临时密钥，避免了您把自己的AccessKey信息编码在函数代码里
    access_key_id = context.credentials.access_key_id
    access_key_secret = context.credentials.access_key_secret
    security_token = context.credentials.security_token
    # 获取自己 mns 的推送地址 Endpoint
    mns_endpoint = os.getenv("MNS_ENDPOINT")
    # 获取自己 mns queue 的名称
    queue_name = os.getenv("MNS_QUEUE_NAME")
    # 创建mns实例
    my_account = Account(mns_endpoint, access_key_id,
                         access_key_secret, security_token)
    # 获取mns实例的一个 queue 对象
    my_queue = my_account.get_queue(queue_name)


def handler(event, context):
    logger = context.getLogger()
    
    logger.info("Send Message To Queue {}".format(queue_name))

    try:
        msg_body = "I am a test message."
        msg = Message(msg_body)
        re_msg = my_queue.send_message(msg)
        return "Send Message Succeed. MessageBody:%s MessageID:%s" % (msg_body, re_msg.message_id)
    except MNSExceptionBase as e:
        if e.type == "QueueNotExist":
            logger.info("Queue '{}' not exist, please create queue before send message.".format(queue_name))
        raise RuntimeError(e)
