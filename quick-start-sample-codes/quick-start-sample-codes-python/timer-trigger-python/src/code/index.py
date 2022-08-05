import json
import logging

logger = logging.getLogger()

def handler(event, context):
    logger.info('event: %s', event)

    # Parse the json
    evt = json.loads(event)
    triggerName = evt["triggerName"]
    triggerTime = evt["triggerTime"]
    payload = evt["payload"]

    logger.info('triggerName: %s', triggerName)
    logger.info("triggerTime: %s", triggerTime)
    logger.info("payload: %s", payload)     

    return 'Timer Payload: ' + payload