package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

var (
	client ali_mns.MNSClient
	topic  ali_mns.AliMNSTopic
)

func initialize(ctx context.Context) {
	mnsEndpoint := os.Getenv("MNS_ENDPOINT")
	topicName := os.Getenv("MNS_TOPIC_NAME")
	fctx, _ := fccontext.FromContext(ctx)
	client = ali_mns.NewAliMNSClient(mnsEndpoint, fctx.Credentials.AccessKeyId, fctx.Credentials.AccessKeySecret)
	topic = ali_mns.NewMNSTopic(topicName, client)
}

func HandleRequest(ctx context.Context) (string, error) {

	//发送一条消息
	msg := ali_mns.MessagePublishRequest{
		MessageBody: "hello topic <\"aliyun-mns-go-sdk\">",
	}
	messageResponse, err := topic.PublishMessage(msg)
	if err != nil {
		return "fail", err
	}
	return fmt.Sprintf("Publish succ, message id: %s, messagebody md5: %s", messageResponse.MessageId, messageResponse.MessageBodyMD5), nil

}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.Start(HandleRequest)
}
