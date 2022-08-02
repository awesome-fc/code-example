package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	var topicEvent events.MnsTopicEvent
	switch mess := event.(type) {
	case string:
		{
			flog.Info("event:", event)
			return fmt.Sprintf("MessageBody:%s", mess), nil
		}
	default:
		result, _ := json.Marshal(mess)
		_ = json.Unmarshal(result, &topicEvent)
		flog.Info("event:", event)
		return fmt.Sprintf("MessageBody:%s", *topicEvent.Message), nil
	}

}

func main() {
	fc.Start(HandleRequest)
}
