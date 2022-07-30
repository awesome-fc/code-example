package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

type StructEvent struct {
	TriggerTime string `json:"triggerTime"`
	TriggerName string `json:"triggerName"`
	Payload     string `json:"payload"`
}

func HandleRequest(ctx context.Context, event StructEvent) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	flog.Info("triggerName: ", event.TriggerName)
	flog.Info("triggerTime: ", event.TriggerTime)
	flog.Info("payload:", event.Payload)
	return fmt.Sprintf("Timer Payload: %s", event.Payload), nil
}

func main() {
	fc.Start(HandleRequest)
}
