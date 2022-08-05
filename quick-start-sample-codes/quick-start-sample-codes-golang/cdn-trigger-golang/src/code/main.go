package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

/*
* 本代码样例主要实现以下功能:
** 打印 event


* This sample code is mainly doing the following things:
** Print event
 */

type CdnEvent struct {
	Events []*CdnEventRecord `json:"events"`
}

type CdnEventRecord struct {
	EventName      *string         `json:"eventName"`
	EventVersion   *string         `json:""eventVersion`
	EventSource    *string         `json:"eventSource"`
	EventTime      *time.Time      `json:"eventTime"`
	Region         *string         `json:"region"`
	TraceId        *string         `json:"traceId"`
	EventParameter *EventParameter `json:"eventParameter"`
	UserIdentity   *UserIdentity   `json:"UserIdentity"`
	Resource       *Resource       `json:"Resource"`
}

type UserIdentity struct {
	AliUid *string `json:"aliUid"`
}

type Resource struct {
	Domain *string `json:"domain"`
}

type EventParameter struct {
	Domain       *string   `json:"domain"`
	CompleteTime *int      `json:"completeTime"`
	ObjectPath   []*string `json:"objectPath"`
	ObjectType   *string   `json:"objectType"`
	TaskId       *int      `json:"taskId"`
	CreateTime   *int      `json:"createTime"`
	EndTime      *int      `json:"endTime"`
	FileSize     *int      `json:"fileSize"`
	StartTime    *int      `json:"startTime"`
	FilePath     *string   `json:"filePath"`
	Status       *string   `json:"status"`
}

func main() {
	fc.Start(HandleRequest)
}

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
func HandleRequest(ctx context.Context, event CdnEvent) (string, error) {
	eventName := *(event.Events[0].EventName)
	domain := *event.Events[0].EventParameter.Domain
	info := ""
	if eventName == "CachedObjectsRefreshed" || eventName == "CachedObjectsPushed" || eventName == "CachedObjectsBlocked" {
		// 上述三个事件可以获取 objectPath
		objectPath := event.Events[0].EventParameter.ObjectPath
		for _, v := range objectPath {
			info += *v + ","
		}
	} else if eventName == "LogFileCreated" {
		info = *event.Events[0].EventParameter.FilePath
	} else if eventName == "CdnDomainStarted" || eventName == "CdnDomainStopped" {
		// 对应业务逻辑
	} else if eventName == "CdnDomainAdded" || eventName == "CdnDomainDeleted" {
		// 对应业务逻辑
	}
	return fmt.Sprintf("eventName:%s, domain:%s, info:%s", eventName, domain, info), nil
}
