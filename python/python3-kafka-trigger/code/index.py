# -*- coding: utf-8 -*-
import logging
import json

def handler(event, context):
  logger = logging.getLogger()
  logger.info('Receive kafka whole message:' + bytes.decode(event))

  # Parse the json
  evt = json.loads(event)
  # Get the first json object from json array 
  evt = evt[0]
  logger.info('message topic:' + evt['topic'])
  logger.info('message value:' + evt['value'])

  return 'kafka whole message:' + bytes.decode(event)