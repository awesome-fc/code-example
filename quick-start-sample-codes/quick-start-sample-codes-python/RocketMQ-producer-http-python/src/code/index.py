import logging
import os
from mq_http_sdk.mq_exception import MQExceptionBase
from mq_http_sdk.mq_producer import *
from mq_http_sdk.mq_client import *


def initialize(context):
    global msg_producer,topic
    access_key_id = context.credentials.access_key_id
    access_key_secret = context.credentials.access_key_secret
    security_token = context.credentials.security_token
    endpoint = os.getenv('ROCKETMQ_ENDPOINT')
    instanceID = os.getenv('INSTANCEID')
    topic = os.getenv('TOPIC')
    mq_client = MQClient(endpoint, access_key_id, access_key_secret, security_token)
    msg_producer = mq_client.get_producer(instanceID, topic)

def handler(environ, start_response):
    logger = logging.getLogger()
    logger.info("Publishi Message To Topic {}".format(topic))
    status = '200 OK'
    response_headers = [('Content-type', 'text/plain')]

    try:
        msg = TopicMessage(
            "hello rocketmq",
            "tag greeting"
        )
        msg.put_property("a", 1)
        msg.set_message_key("MessageKey")
        re_msg = msg_producer.publish_message(msg)
        start_response(status, response_headers)
        return "Publish Message Succeed. MessageID:%s , BodyMD5 : %s" % (re_msg.message_id, re_msg.message_body_md5)
    except MQExceptionBase as e:
        if e.type == "TopicNotExist":
            logger.info("Topic Not Exist")
        raise RuntimeError(e)