package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func HandleRequest(ctx context.Context, event events.MnsQueueEvent) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		panic("parse context fail")
	}
	logger := fctx.GetLogger()
	mes, err := json.Marshal(event)
	if err != nil {
		logger.Errorf("json.Marshal mns message fail:%v", err)
	} else {
		logger.Info("receive mns message:", string(mes))
	}
	return fmt.Sprintf("MessageBody:%s", *event.Data.MessageBody), nil
}

func main() {
	fc.Start(HandleRequest)
}
