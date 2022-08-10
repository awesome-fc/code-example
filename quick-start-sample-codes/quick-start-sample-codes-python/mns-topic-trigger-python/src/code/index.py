import logging
import json


def handler(event, context):
    logger = logging.getLogger()
    logger.info('Receive mns topic whole message:{}'.format(event))
    # Parse the json
    try:
        eventObect = json.loads(event)
    except Exception:
        return "the event format is STREAM and mns topic message content is:{}".format(event)
    else:
        if 'Message' in eventObect:
            return "the event format is JSON and mns topic message content is:{}".format(eventObect['Message'])
        return "the event format is STREAM and mns topic message content is:{}".format(event)
