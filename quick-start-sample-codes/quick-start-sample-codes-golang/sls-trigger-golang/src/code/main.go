package main

/*
* 本代码样例主要实现以下功能:
*   1. 从 event 中解析出 SLS 事件触发相关信息
*   2. 根据以上获取的信息，初始化 SLS 客户端
*   3. 从源日志仓库中获取实时日志数据
*
*
* This sample code is mainly doing the following things:
*   1. Get SLS processing related information from event
*   2. Initiate SLS client
*   3. Pull logs from source log store
* */

import (
	"context"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func main() {
	fc.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event events.SlsEvent) (string, error) {
	// 获取密钥信息，执行前，确保函数所在的服务配置了角色信息，并且角色需要拥有 AliyunLogFullAccess权限
	fctx, _ := fccontext.FromContext(ctx)
	creds := fctx.Credentials
	logger := fctx.GetLogger()

	// 从 event.source 中获取日志项目名称、日志仓库名称以及日志服务访问 endpoint
	// Get the name of log project, the name of log store and the endpoint of sls from event.source
	projectName := *event.Source.ProjectName
	logstoreName := *event.Source.LogstoreName
	endpoint := *event.Source.Endpoint
	beginCursor := *event.Source.BeginCursor
	endCursor := *event.Source.EndCursor
	shardID := *event.Source.ShardID

	// 初始化 sls 客户端
	// Initialize client of sls
	client := sls.CreateNormalInterface(endpoint, creds.AccessKeyId, creds.AccessKeySecret, creds.SecurityToken)

	// 从源日志库中读取日志
	// Read data from source logstore
	for {
		logGroupList, nextCursor, err := client.PullLogs(projectName, logstoreName, shardID, beginCursor, endCursor, 100)
		if err != nil {
			logger.Errorf("pull logs failed due to: %+v", err)
		}
		if len(logGroupList.LogGroups) == 0 {
			break
		}
		beginCursor = nextCursor

		logger.Infof("Get %d log group from %s", len(logGroupList.LogGroups), logstoreName)
		for _, logGroup := range logGroupList.LogGroups {
			logger.Infof("Topic: %s, Source: %s", *logGroup.Topic, *logGroup.Source)
			for _, log := range logGroup.Logs {
				logger.Infof("Log time: %d", *log.Time)
				for _, content := range log.Contents {
					logger.Infof("Log content key: %s, value: %s", *content.Key, *content.Value)
				}
			}
		}
	}

	return "success", nil
}
