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
	GroupId          string `json:"group.id"`
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
	bootstrapServers = os.Getenv("bootstrap_servers")
	topicName = os.Getenv("topic_name")

	fctx, _ := fccontext.FromContext(ctx)
	fctx.GetLogger().Infof("Initializing the kafka config\n")

	var kafkaconf = &kafka.ConfigMap{
		"api.version.request": "true",
	}
	kafkaconf.SetKey("bootstrap.servers", bootstrapServers)
	kafkaconf.SetKey("security.protocol", "plaintext")

	producer, _ = kafka.NewProducer(kafkaconf)

}

func HandleRequest(ctx context.Context, event StructEvent) (string, error) {
	topic := topicName

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(event.Key),
	}, nil)
	fctx, _ := fccontext.FromContext(ctx)
	fctx.GetLogger().Infof("sending the message to kafka: %s!", event.Key)

	producer.Flush(1000)
	return fmt.Sprintf("Finish sending the message to kafka: %s!", event.Key), nil
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.Start(HandleRequest)
}
