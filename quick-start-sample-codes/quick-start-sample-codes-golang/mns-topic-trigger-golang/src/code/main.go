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
	flog.Info("Receive mns topic whole message:", event)
	switch mess := event.(type) {
	// the format of event is STREAM
	case string:
		{
			return fmt.Sprintf("the event format is STREAM and message is:%s", mess), nil
		}
	// the format of event is JSON
	default:
		result, _ := json.Marshal(mess)
		_ = json.Unmarshal(result, &topicEvent)
		return fmt.Sprintf("the event format is JSON and MessageBody is:%s", *topicEvent.Message), nil
	}
}

func main() {
	fc.Start(HandleRequest)
}
