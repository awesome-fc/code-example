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
	"fmt"
	"os"
	"strconv"

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

	// 从 event 中获取 cursorTime，该字段表示本次函数调用包括的数据中，最后一条日志到达日志服务的服务器端的 unix_timestamp
	// Get cursorTime from event, where cursorTime indicates that in the data of the invocation, the unix timestamp of the last log arrived at log store
	cursorTime := *event.CursorTime

	// 从 event.source 中获取日志项目名称、日志仓库名称以及日志服务访问 endpoint
	// Get the name of log project, the name of log store and the endpoint of sls from event.source
	projectName := *event.Source.ProjectName
	logstoreName := *event.Source.LogstoreName
	endpoint := *event.Source.Endpoint

	// 初始化 sls 客户端
	// Initialize client of sls
	client := sls.CreateNormalInterface(endpoint, creds.AccessKeyId, creds.AccessKeySecret, creds.SecurityToken)

	// 从环境变量中获取触发时间间隔，该环境变量可在 s.yml 中配置
	// Get interval of trigger from environment variables, which was configured via s.yaml
	triggerInterval, err := strconv.Atoi(os.Getenv("triggerInterval"))
	if err != nil {
		fmt.Printf("converse triggerInterval from string to int fail due to: %+v", err)
		return "", err
	}

	// 从源日志库中读取日志
	// Read data from source logstore
	toTime := int64(cursorTime)
	fromTime := int64(cursorTime - triggerInterval)
	getLogsRes, err := client.GetLogs(projectName, logstoreName, "", fromTime, toTime, "", 1000, 0, true)
	if err != nil {
		fmt.Printf("get logs failed due to: %+v", err)
		return "", err
	}
	fmt.Printf("Read log data count: %d\n", getLogsRes.Count)
	fmt.Printf("log Contents: %s\n", getLogsRes.Contents)
	fmt.Printf("log Logs: %s\n", getLogsRes.Logs)
	fmt.Printf("log Progress: %s\n", getLogsRes.Progress)



	return "success", nil
}