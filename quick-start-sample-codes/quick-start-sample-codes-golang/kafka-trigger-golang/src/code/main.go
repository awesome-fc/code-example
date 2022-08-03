package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

// Define the kafka trigger event struct to be received
type StructEvent struct {
	Data            KafkaData `json:"data"`
	Id              string    `json:"id"`
	Sourse          string    `json:"sourse"`
	Specversion     string    `json:"specversion"`
	Type            string    `json:"type"`
	Datacontenttype string    `json:"datacontenttype"`
	Time            string    `json:"time"`
	Subject         string    `json:"subject"`
	Aliyunaccountid string    `json:"aliyunaccountid"`
}
type KafkaData struct {
	Topic     string     `json:"topic"`
	Partition int        `json:"partition"`
	Offset    int        `json:"offset"`
	Timestamp int        `json:"timestamp"`
	Headers   DataHeader `json:"headers"`
	Value     string     `json:"value"`
}

type DataHeader struct {
	Headers    []string `json:"headers"`
	IsReadOnly bool     `json:"isReadOnly"`
}

func HandleRequest(ctx context.Context, event []string) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()

	eventString := event[0]
	var evt StructEvent
	json.Unmarshal([]byte(eventString), &evt)

	flog.Info("kafka event:", event)
	// The trigger event data is in the `Data` json object from the json array
	flog.Info("kafka topic:", evt.Data.Topic)
	flog.Info("kafka messgae:", evt.Data.Value)

	return fmt.Sprintf("Receive Kafka Messgae Value: %s", evt.Data.Value), nil
}

func main() {
	fc.Start(HandleRequest)
}
