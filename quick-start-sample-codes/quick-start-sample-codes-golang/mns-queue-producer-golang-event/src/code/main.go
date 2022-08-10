package main

import (
	"context"
	"fmt"
	"os"

	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

var (
	client ali_mns.MNSClient
	queue  ali_mns.AliMNSQueue
)

func initialize(ctx context.Context) {
	mnsEndpoint := os.Getenv("MNS_ENDPOINT")
	queueName := os.Getenv("MNS_QUEUE_NAME")
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		panic("parse context fail")
	}
	client = ali_mns.NewAliMNSClient(mnsEndpoint, fctx.Credentials.AccessKeyId, fctx.Credentials.AccessKeySecret)
	queue = ali_mns.NewMNSQueue(queueName, client)
}

func HandleRequest(ctx context.Context) (string, error) {
	//发送一条消息
	msg := ali_mns.MessageSendRequest{
		MessageBody:  "hello <\"aliyun-mns-go-sdk\">",
		DelaySeconds: 0,
		Priority:     8}
	ret, err := queue.SendMessage(msg)
	if err != nil {
		return "fail", err
	}
	return fmt.Sprintf("Send succ, message id: %s, messagebody md5: %s", ret.MessageId, ret.MessageBodyMD5), nil
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.Start(HandleRequest)
}
