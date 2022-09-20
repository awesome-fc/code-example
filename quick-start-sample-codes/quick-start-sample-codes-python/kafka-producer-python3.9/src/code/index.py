from confluent_kafka import Producer
import logging
import sys
import os

logger = logging.getLogger()

def initialize(context):
    global p, TOPIC_NAME

    """ Get the environment variables """
    BOOTSTRAP_SERVERS = os.getenv("BOOTSTRAP_SERVERS")
    TOPIC_NAME = os.getenv("TOPIC_NAME")

    p = Producer({'bootstrap.servers': BOOTSTRAP_SERVERS})

def deliveryReport(err, msg):
    """ Called once for each message produced to indicate delivery result.
        Triggered by poll() or flush(). """
    if err is not None:
        logger.info('Message delivery failed: {}'.format(err))
        raise Exception('Message delivery failed: {}'.format(err))
    else:
        logger.info('Message delivered to {} [{}]'.format(msg.topic(), msg.partition()))

def handler(event, context):
    """ Produce messages to topic (asynchronously) """
    p.produce(TOPIC_NAME, str(event).encode('utf-8'), callback=deliveryReport)
    p.poll(0)

    """ Flush the internel queue, wait for message deliveries before return """
    p.flush()

    return "finish sending message: " + str(event)
