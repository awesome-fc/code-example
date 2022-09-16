# -*- coding: utf-8 -*-
import logging
import os
from tablestore import OTSClient


logger = logging.getLogger()

client: OTSClient = None

def initialize(context):
    # 在initialize回调中创建客户端，可以实现在整个函数实例生命周期内复用该客户端
    global client
    creds = context.credentials
    access_key, access_secret, sts_token = creds.access_key_id, creds.access_key_secret, creds.security_token
    client = OTSClient(os.getenv("ENDPOINT"), access_key, access_secret, os.getenv("INSTANCE_NAME"), sts_token=sts_token)

def parse_query(query_str):
    ret = {}
    for q in query_str.split("&"):
        k, v = q.split("=")[0:2]
        ret[k] = v
    return ret

def handler(environ, start_response):
    # 本示例所用表格存储的主键包含两个主键列：region 和 id
    query_dict = parse_query(environ['QUERY_STRING'])
    primary_key = [('region', query_dict['region']), ('id', int(query_dict['id']))]
    # 要返回的属性列，如果columns_to_get为[]，则返回所有属性列
    columns_to_get = []
    _, return_row, _ = client.get_row(table_name=os.getenv("TABLE_NAME"), primary_key=primary_key, columns_to_get=columns_to_get)
    status = '200 OK'
    response_headers = [('Content-type', 'application/json')]
    start_response(status, response_headers)
    return [str(return_row.attribute_columns)]



