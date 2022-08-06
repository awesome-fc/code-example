"""
本代码样例主要实现以下功能:
* 从 event 中解析出 SLS 事件触发相关信息
* 根据以上获取的信息，初始化 SLS 客户端
* 从源 log store 获取实时日志数据


This sample code is mainly doing the following things:
* Get SLS processing related information from event
* Initiate SLS client
* Pull logs from source log store

"""
#!/usr/bin/env python
# -*- coding: utf-8 -*-

import logging
import json
import os
from aliyun.log import LogClient


logger = logging.getLogger()


def handler(event, context):

    # 可以通过 context.credentials 获取密钥信息
    # Access keys can be fetched through context.credentials
    print("The content in context entity is: \n")
    print(context)
    creds = context.credentials
    access_key_id = creds.access_key_id
    access_key_secret = creds.access_key_secret
    security_token = creds.security_token

    # 解析 event 参数至 object 格式
    # parse event in object
    event_obj = json.loads(event.decode())
    print("The content in event entity is: \n")
    print(event_obj)

    # 从 event.source 中获取日志项目名称、日志仓库名称、日志服务访问 endpoint、日志起始游标、日志终点游标以及分区 id
    # Get the name of log project, the name of log store, the endpoint of sls, begin cursor, end cursor and shardId from event.source
    source = event_obj['source']
    log_project = source['projectName']
    log_store = source['logstoreName']
    endpoint = source['endpoint']
    begin_cursor = source['beginCursor']
    end_cursor = source['endCursor']
    shard_id = source['shardId']

    # 初始化 sls 客户端
    # Initialize client of sls
    client = LogClient(endpoint=endpoint, accessKeyId=access_key_id, accessKey=access_key_secret, securityToken=security_token)

    # 基于日志的游标从源日志库中读取日志，本示例中的游标范围包含了触发本次执行的所有日志内容
    # Read data from source logstore within cursor: [begin_cursor, end_cursor) in the example, which contains all the logs trigger the invocation
    while True:
      response = client.pull_logs(project_name=log_project, logstore_name=log_store,
                                shard_id=shard_id, cursor=begin_cursor, count=100,
                                end_cursor=end_cursor, compress=False)
      log_group_cnt = response.get_loggroup_count()
      if log_group_cnt == 0:
        break
      logger.info("get %d log group from %s" % (log_group_cnt, log_store))
      logger.info(response.get_loggroup_list())

      begin_cursor = response.get_next_cursor()

    return 'success'