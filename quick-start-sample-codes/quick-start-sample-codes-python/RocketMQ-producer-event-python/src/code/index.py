import logging
import os
from mq_http_sdk.mq_exception import MQExceptionBase
from mq_http_sdk.mq_producer import *
from mq_http_sdk.mq_client import *

logger = logging.getLogger()


def initialize(context):
    global msg_producer
    access_key_id = context.credentials.access_key_id
    access_key_secret = context.credentials.access_key_secret
    security_token = context.credentials.security_token
    endpoint = os.getenv('ROCKETMQ_ENDPOINT')
    instanceID = os.getenv('INSTANCEID')
    topic = os.getenv('TOPIC')
    mq_client = MQClient(endpoint, access_key_id, access_key_secret, security_token)
    msg_producer = mq_client.get_producer(instanceID, topic)


def handler(event, context):
    try:
        msg = TopicMessage(
            "hello rocketmq",
            "tag greeting"
        )
        msg.put_property("a", 1)
        msg.set_message_key("MessageKey")
        re_msg = msg_producer.publish_message(msg)
        return "Publish Message Succeed. MessageID:%s , BodyMD5 : %s" % (re_msg.message_id, re_msg.message_body_md5)
    except MQExceptionBase as e:
        if e.type == "TopicNotExist":
            logger.info("Topic not exist. please create it.")
            sys.exit(1)
        logger.info("Publish Message Fail. Exception:%s" % e)
