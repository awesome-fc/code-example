package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConfig struct {
	Topic            string `json:"topic"`
	BootstrapServers string `json:"bootstrap.servers"`
	SecurityProtocol string `json:"security.protocol"`
}

type StructEvent struct {
	Key string `json:"key"`
}

var producer *kafka.Producer
var bootstrapServers string
var topicName string

func initialize(ctx context.Context) {
	// Get the environment variables
	bootstrapServers = os.Getenv("BOOTSTRAP_SERVERS")
	topicName = os.Getenv("TOPIC_NAME")

	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		panic("Get fccontext fail.")
	}
	fctx.GetLogger().Infof("Initializing the kafka config")

	var kafkaconf = &kafka.ConfigMap{
		"api.version.request": "true",
	}
	kafkaconf.SetKey("bootstrap.servers", bootstrapServers)
	kafkaconf.SetKey("security.protocol", "plaintext")

	var err error
	producer, err = kafka.NewProducer(kafkaconf)
	if err != nil {
		panic(err)
	}
}

func HandleRequest(ctx context.Context, event StructEvent) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return "Get fccontext fail.", nil
	}
	fctx.GetLogger().Infof("sending the message to kafka: %s!", event.Key)

	// Produce messages to topic (synchronously)
	delivery_chan := make(chan kafka.Event, 1)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Value:          []byte(event.Key)}, delivery_chan)

	e := <-delivery_chan
	m := e.(*kafka.Message)

	// Capture the delivery report
	if m.TopicPartition.Error != nil {
		return "Something wrong with TopicPartition", m.TopicPartition.Error
	} else {
		fctx.GetLogger().Infof("Delivered message to topic %s [%d] at offset %v",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(delivery_chan)

	return fmt.Sprintf("Finish sending the message to kafka: %s!", event.Key), nil
}

func preStop(ctx context.Context) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return
	}
	fctx.GetLogger().Infof("preStop hook start.")
	if producer != nil {
		producer.Close()
	}

	fctx.GetLogger().Infof("preStop hook finish.")
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.RegisterPreStopFunction(preStop)
	fc.Start(HandleRequest)
}
