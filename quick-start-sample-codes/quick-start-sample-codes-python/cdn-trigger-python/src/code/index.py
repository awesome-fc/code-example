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
# 各 event 示例见文档：https://help.aliyun.com/document_detail/75123.html，event结构如下所示：

#  {  "events": [
#        {
#           "eventName": "***",
#           "eventVersion": "***",
#           "eventSource": "***",
#           "region": "***",
#           "eventTime": "***",
#           "traceId": "***",
#           "resource": {
#                "domain": "***"
#           },
#           "eventParameter": {

#           },
#           "userIdentity": {
#                "aliUid": "***"
#           }
#        }
#     ]
#  }
def handler(event, context):
    evt = json.loads(event)
    eventObj = evt["events"][0]
    eventName = eventObj['eventName']
    info = ""
    eventParam = eventObj['eventParameter']
    domain = eventParam['domain']
    if eventName == "CachedObjectsRefreshed" or eventName == "CachedObjectsPushed" or eventName == "CachedObjectsBlocked":
        objPathList = eventParam['objectPath']
        info = ",".join(objPathList)
    elif eventName == "LogFileCreated":
        info = eventParam['filePath']
    elif eventName == "CdnDomainStarted" or eventName == "CdnDomainStopped":
        # 对应业务逻辑
        pass
    elif eventName == "CdnDomainAdded" or eventName == "CdnDomainDeleted":
        # 对应业务逻辑
        pass
    
    return f"eventName:{eventName}, domain: {domain}, info: {info}"
