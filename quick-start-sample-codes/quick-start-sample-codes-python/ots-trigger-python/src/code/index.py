"""
本代码样例主要实现以下功能:
* 打印 event 信息


This sample code is mainly doing the following things:
* print event

"""

# -*- coding: utf-8 -*-
import logging
import cbor


logger = logging.getLogger("ots-sample")
# event 示例见文档：https://help.aliyun.com/document_detail/169672.html，event结构如下所示：
# {
#     "Version": "Sync-v1",
#     "Records": [
#         {
#             "Type": "PutRow",
#             "Info": {
#                 "Timestamp": 1506416585740836
#             },
#             "PrimaryKey": [
#                 {
#                     "ColumnName": "pk_0",
#                     "Value": 1506416585881590900
#                 },
#                 {
#                     "ColumnName": "pk_1",
#                     "Value": "2017-09-26 17:03:05.8815909 +0800 CST"
#                 },
#                 {
#                     "ColumnName": "pk_2",
#                     "Value": 1506416585741000
#                 }
#             ],
#             "Columns": [
#                 {
#                     "Type": "Put",
#                     "ColumnName": "attr_0",
#                     "Value": "hello_table_store",
#                     "Timestamp": 1506416585741
#                 },
#                 {
#                     "Type": "Put",
#                     "ColumnName": "attr_1",
#                     "Value": 1506416585881590900,
#                     "Timestamp": 1506416585741
#                 }
#             ]
#         }
#     ]
# }

def handler(event, context):
    logger = logging.getLogger()
    logger.info("Begin to handle event")
    records = cbor.loads(event)
    logger.info(f"Record:{records}")
    for record in records['Records']:
        primary_values = record['PrimaryKey']
        attribute_values = record['Columns']
    return 'OK'
