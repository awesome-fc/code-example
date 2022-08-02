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
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	mes, _ := json.Marshal(event)
	flog.Info("event:", string(mes))
	return fmt.Sprintf("MessageBody:%s", *event.Data.MessageBody), nil
}

func main() {
	fc.Start(HandleRequest)
}
