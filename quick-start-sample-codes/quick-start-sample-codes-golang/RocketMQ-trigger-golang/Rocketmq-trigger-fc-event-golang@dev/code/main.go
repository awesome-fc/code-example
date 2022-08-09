package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	mes, _ := json.Marshal(event)
	flog.Info("event: ", string(mes))
	sendData := gojsonq.New().FromString(string(mes)).Find("data.body")
	return fmt.Sprintf("queue data: %v ", sendData), nil
}

func main() {
	fc.Start(HandleRequest)
}
