# -*- coding: utf-8 -*-
import logging
import json

def handler(event, context):
  logger = logging.getLogger()
  logger.info('Receive rabbitMQ whole message:' + bytes.decode(event))

  # Parse the json
  eventObject = json.loads(event)
  # Get each json object from json array
  for evt in eventObject:
    # Parse the json inside
    evt = json.loads(evt)

    logger.info('message body:' + evt['data']['body'])
    logger.info('message id:' + evt['data']['props']['messageId'])

  return 'RabbitMQ Trigger Event:' + bytes.decode(event)
