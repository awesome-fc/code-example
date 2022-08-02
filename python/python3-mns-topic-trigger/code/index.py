import logging


def handler(event, context):
    logger = logging.getLogger()
    logger.info(event)
    return "mns_topic trigger event = {}".format(event)
