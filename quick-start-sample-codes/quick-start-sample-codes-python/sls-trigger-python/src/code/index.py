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
from aliyun.log import LogClient, PutLogsRequest, LogItem, GetLogsRequest, GetLogsResponse


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
    
    # 从 event 中获取 cursorTime，该字段表示本次函数调用包括的数据中，最后一条日志到达日志服务的服务器端的 unix_timestamp
    # Get cursorTime from event, where cursorTime indicates that in the data of the invocation, the unix timestamp of the last log arrived at log store
    cursor_time = event_obj['cursorTime']

    # 从 event.source 中获取日志项目名称、日志仓库名称以及日志服务访问 endpoint
    # Get the name of log project, the name of log store and the endpoint of sls from event.source
    source = event_obj['source']
    log_project = source['projectName']
    log_store = source['logstoreName']
    endpoint = source['endpoint']

    # 从环境变量中获取触发时间间隔，该环境变量可在 s.yml 中配置
    # Get interval of trigger from environment variables, which was configured via s.yaml
    trigger_interval = int(os.environ.get('triggerInterval'))

    # 初始化 sls 客户端
    # Initialize client of sls
    client = LogClient(endpoint=endpoint, accessKeyId=access_key_id, accessKey=access_key_secret, securityToken=security_token)

    # 从源日志库中读取日志
    # Read data from source logstore
    to_time = cursor_time
    from_time = to_time - trigger_interval

    request = GetLogsRequest(project=log_project, logstore=log_store, fromTime=from_time, toTime=to_time, query='')
    response = client.get_logs(request)

    for log in response.get_logs():
        logger.info(log.contents.items())

    return 'success'