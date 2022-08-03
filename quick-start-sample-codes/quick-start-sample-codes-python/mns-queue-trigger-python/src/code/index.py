import logging


def handler(event, context):
    logger = logging.getLogger()
    logger.info(event)
    return "mns_queue trigger event = {}".format(event)
