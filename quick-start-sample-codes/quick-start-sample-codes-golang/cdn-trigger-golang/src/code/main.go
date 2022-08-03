package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

/*
* 本代码样例主要实现以下功能:
** 打印 event


* This sample code is mainly doing the following things:
** Print event
 */

func main() {
	fc.Start(HandleRequest)
}

func MapToJsonStr(data map[string]interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func HandleRequest(ctx context.Context, event map[string]interface{}) (string, error) {
	fmt.Println(event)
	return MapToJsonStr(event), nil
}
