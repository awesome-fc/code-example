# -*- coding: utf-8 -*-
import logging
import os
from tablestore import OTSClient


logger = logging.getLogger()

client: OTSClient = None

def initialize(context):
    # 在initialize回调中创建客户端，可以实现在整个函数实例生命周期内复用该客户端
    global client
    client = OTSClient(os.getenv("ENDPOINT"), os.getenv("ACCESS_KEY"), os.getenv("ACCESS_KEY_SECRET"), os.getenv("INSTANCE_NAME"))


def handler(event, context):
    # 本示例所用表格存储的主键包含两个主键列：region 和 id
    primary_key = [('region', "abc"), ('id', 1)]
    # 要返回的属性列，如果columns_to_get为[]，则返回所有属性列
    columns_to_get = []
    _, return_row, _ = client.get_row(table_name="fc_test", primary_key=primary_key, columns_to_get=columns_to_get)
    return str(return_row.attribute_columns) 



