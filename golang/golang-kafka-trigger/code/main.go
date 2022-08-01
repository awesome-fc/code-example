package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

// Define the kafka trigger event struct to be received
type StructEvent struct {
	Key          string `json:"key"`
	Value        string `json:"value"`
	Topic        string `json:"topic"`
	Offset       int    `json:"offset"`
	OverflowFlag bool   `json:"overflowFlag"`
	Partition    int    `json:"partition"`
	Timestamp    int    `json:"timestamp"`
	ValueSize    int    `json:"valueSize"`
}

func HandleRequest(ctx context.Context, event []StructEvent) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()

	flog.Info("kafka event:", event)
	// The trigger event is the first json object from the json array
	flog.Info("kafka topic:", event[0].Topic)
	flog.Info("kafka messgae:", event[0].Value)

	return fmt.Sprintf("Receive Kafka Messgae Value: %s", event[0].Value), nil
}

func main() {
	fc.Start(HandleRequest)
}
