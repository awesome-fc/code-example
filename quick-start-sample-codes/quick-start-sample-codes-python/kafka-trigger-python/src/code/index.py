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
  # Parse the json inside
  evt = json.loads(evt)
  
  logger.info('message topic:' + evt['data']['topic'])
  logger.info('message value:' + evt['data']['value'])

  return 'Kafka trigger data message:' + evt['data']['value']
