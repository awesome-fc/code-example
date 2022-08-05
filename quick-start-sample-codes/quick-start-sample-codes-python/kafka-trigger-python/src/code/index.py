# -*- coding: utf-8 -*-
import logging
import json

def handler(event, context):
  logger = logging.getLogger()
  logger.info('Receive kafka whole message:' + bytes.decode(event))

  # Parse the json
  eventObject = json.loads(event)
  # Get each json object from json array
  for evt in eventObject:
    # Parse the json inside
    evt = json.loads(evt)

    logger.info('message topic:' + evt['data']['topic'])
    logger.info('message value:' + evt['data']['value'])

  return 'Kafka Trigger Event:' + bytes.decode(event)
