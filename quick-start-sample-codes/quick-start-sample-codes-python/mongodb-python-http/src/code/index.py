# -*- coding: utf-8 -*-
import logging
import pymongo
import os
from urllib import parse

logger = logging.getLogger()

client = None

def initialize(context):
    # 在initialize回调中创建客户端，可以实现在整个函数实例生命周期内复用该客户端
    global client
    client = pymongo.MongoClient(os.environ['MONGO_URL'])


def pre_stop(context):
    if client != None:
        client.close()

def parse_query(query_str):
    ret = {}
    for q in query_str.split("&"):
        k, v = q.split("=")[0:2]
        ret[k] = v
    return ret


def handler(environ, start_response):
    query_dict = parse_query(environ['QUERY_STRING'])
    collection = client[os.environ['MONGO_DATABASE']]['users']
    res = collection.find_one({
            "name": parse.unquote(query_dict['name'])
        }
    )
    status = '200 OK'
    response_headers = [('Content-type', 'application/json')]
    start_response(status, response_headers)
    return [str(res)]


