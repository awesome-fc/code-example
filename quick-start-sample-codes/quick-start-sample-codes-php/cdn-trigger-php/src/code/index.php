<?php
/*
* 本代码样例主要实现以下功能:
*   1. 打印 event 信息

* This sample code is mainly doing the following things:
*  1. print event
* */

// 各 event 示例见文档：https://help.aliyun.com/document_detail/75123.html，event结构如下所示：
//
// {  "events": [
//       {
//          "eventName": "***",
//          "eventVersion": "***",
//          "eventSource": "***",
//          "region": "***",
//          "eventTime": "***",
//          "traceId": "***",
//          "resource": {
//               "domain": "***"
//          },
//          "eventParameter": {
//
//          },
//          "userIdentity": {
//               "aliUid": "***"
//          }
//       }
//    ]
// }

function handler($event, $context){
    $eventJson = json_decode($event, true);
    $eventName = $eventJson['events'][0]['eventName'];
    $domain = $eventJson['events'][0]['eventParameter']['domain'];
    $eventParam = $eventJson['events'][0]['eventParameter'];
    $info = "";
    if($eventName == "CachedObjectsRefreshed" || $eventName == "CachedObjectsPushed" || $eventName == "CachedObjectsBlocked"){
        $objPathList = $eventParam['objectPath'];
        foreach($objPathList as $v){
            $info .= $v . ",";
        }
    }else if($eventName == "LogFileCreated"){
        $info = $eventParam['filePath'];
    }else if($eventName == "CdnDomainStarted" || $eventName == "CdnDomainStopped"){
        // 对应业务逻辑
    }else if($eventName == "CdnDomainAdded" || $eventName == "CdnDomainDeleted"){
        // 对应业务逻辑
    }

    $res = "eventName: " . $eventName . ", domain: " . $domain . ",  info: " . $info;
    return $res;
}