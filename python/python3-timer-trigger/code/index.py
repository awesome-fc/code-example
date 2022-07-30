import json
import logging

logger = logging.getLogger()

def handler(event, context):
    logger.info('event: %s', event)

    evt = json.loads(event)
    triggerName = evt["triggerName"]
    triggerTime = evt["triggerTime"]
    message = evt["payload"]

    logger.info('triggerName: %s', triggerName)
    logger.info("triggerTime = %s", triggerTime)
    logger.info("payload = %s", message)     

    return 'Timer Payload:' + message