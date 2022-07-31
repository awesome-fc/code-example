# -*- coding: utf-8 -*-
import logging
import json

# To enable the initializer feature
# please implement the initializer function as below:
# def initializer(context):
#   logger = logging.getLogger()
#   logger.info('initializing')

def handler(event, context):
  logger = logging.getLogger()
  logger.info('Receive kafka whole message:' + bytes.decode(event))

  evt = json.loads(event)
  evt = evt[0]
  logger.info('message topic:' + evt['topic'])
  logger.info('message value:' + evt['value'])

  return 'kafka whole message:' + bytes.decode(event)