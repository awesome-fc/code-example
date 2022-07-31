from confluent_kafka import Producer
import logging
import sys
import os

logger = logging.getLogger()

def initialize(context):
    global p, topic_name
    bootstrap_servers = os.getenv("bootstrap_servers")
    topic_name = os.getenv("topic_name")
    p = Producer({'bootstrap.servers': bootstrap_servers})

def delivery_report(err, msg):
    """ Called once for each message produced to indicate delivery result.
        Triggered by poll() or flush(). """
    if err is not None:
        logger.info('Message delivery failed: {}'.format(err))
    else:
        logger.info('Message delivered to {} [{}]'.format(msg.topic(), msg.partition()))


def handler(event, context):
    """异步发送消息"""
    p.produce(topic_name, str(event).encode('utf-8'), callback=delivery_report)
    p.poll(0)

    """在函数结束时, 调用flush"""
    p.flush()
    return "finish sending message: " + str(event)