# -*- coding: utf-8 -*-
import logging
import pymongo
import os

logger = logging.getLogger()

client = None

def initialize(context):
    # 在initialize回调中创建客户端，可以实现在整个函数实例生命周期内复用该客户端
    global client
    client = pymongo.MongoClient(os.environ['MONGO_URL'])


def pre_stop(context):
    if client != None:
        client.close()


def handler(event, context):
    collection = client[os.environ['MONGO_DATABASE']]['users']
    res = collection.find_one({"name": "张三"})
    return str(res)


