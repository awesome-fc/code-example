"""
本代码样例主要实现以下功能:
* 从 event 中解析事件列表。
* 将解析的数据插入到 ODPS 数据表中。

This sample code is mainly doing the following things:
* Parse the event list from event param.
* Insert payload into the ODPS table.

"""

# -*- coding: utf-8 -*-

import os
import json
import logging

from odps import ODPS

def handler(event, context):
    logger = logging.getLogger()
    try:
        # 初始化 odps 客户端
        # Initialize odps client.
        env = os.environ
        odps_client = ODPS(env.get('accountAccessKeyID'), env.get('accountAccessKeySecret'), env.get('odpsProject'), env.get('odpsEndpoint'))

        # 解析请求的 payload 数据, 并将 payload 数据插入到 odps 数据表中
        # Parse request payload, and insert the data into the odps table.
        payload = json.loads(event)
        odps_client.write_table(os.environ.get('odpsTableName'), payload)
    except Exception as e:
        logger.error("write odps table failed", e)
        raise json.dumps({"success": False, "error_message": str(e)})

    return json.dumps({"success": True, "error_message": ""})