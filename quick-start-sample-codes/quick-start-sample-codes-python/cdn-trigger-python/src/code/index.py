"""
本代码样例主要实现以下功能:
* 打印 event 信息


This sample code is mainly doing the following things:
* print event

"""

# -*- coding: utf-8 -*-
import logging
import json


logger = logging.getLogger("cdn-sample")

def handler(event, context):
    evt = json.loads(event)
    logger.info(evt)
    return evt
