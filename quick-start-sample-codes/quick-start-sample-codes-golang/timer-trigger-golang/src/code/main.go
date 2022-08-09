package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

// Define the timer trigger event struct
type StructEvent struct {
	TriggerTime string
	TriggerName string
	Payload     string
}

func HandleRequest(ctx context.Context, event StructEvent) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return "Get fccontext fail.", nil
	}
	flog := fctx.GetLogger()

	flog.Info("triggerName: ", event.TriggerName)
	flog.Info("triggerTime: ", event.TriggerTime)
	flog.Info("payload:", event.Payload)

	return fmt.Sprintf("Timer Payload: %s", event.Payload), nil
}

func main() {
	fc.Start(HandleRequest)
}
