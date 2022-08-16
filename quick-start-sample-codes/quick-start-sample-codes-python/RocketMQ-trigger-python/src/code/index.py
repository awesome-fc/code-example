import logging


def handler(event, context):
    logger = logging.getLogger()
    logger.info(event)
    return "RocketMQ trigger event = {}".format(event)