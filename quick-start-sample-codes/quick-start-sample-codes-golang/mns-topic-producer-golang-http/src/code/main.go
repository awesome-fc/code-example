package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

var (
	client ali_mns.MNSClient
	topic  ali_mns.AliMNSTopic
)

func handleError(w http.ResponseWriter, statusCode int, err error) error {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(err.Error()))
	return err
}

func initialize(ctx context.Context) {
	mnsEndpoint := os.Getenv("MNS_ENDPOINT")
	topicName := os.Getenv("MNS_TOPIC_NAME")
	fctx, _ := fccontext.FromContext(ctx)
	client = ali_mns.NewAliMNSClient(mnsEndpoint, fctx.Credentials.AccessKeyId, fctx.Credentials.AccessKeySecret)
	topic = ali_mns.NewMNSTopic(topicName, client)
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	//发送一条消息
	msg := ali_mns.MessagePublishRequest{
		MessageBody: "hello topic <\"aliyun-mns-go-sdk\">",
	}
	messageResponse, err := topic.PublishMessage(msg)
	if err != nil {
		return handleError(w, http.StatusBadRequest, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	_, err = w.Write([]byte(fmt.Sprintf("Publish succ, message id: %s, messagebody md5: %s", messageResponse.MessageId, messageResponse.MessageBodyMD5)))
	if err != nil {
		return handleError(w, http.StatusServiceUnavailable, err)
	}
	return nil
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.StartHttp(HandleHttpRequest)
}
