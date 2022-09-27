package main

import (
	"context"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
	"net/http"
	"os"
	"time"
)

func handleError(w http.ResponseWriter, statusCode int, err error) error {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(err.Error()))
	return err
}
func HandleRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		panic("parse context fail")
	}
	accessKey := fctx.Credentials.AccessKeyId
	secretKey := fctx.Credentials.AccessKeySecret
	stsToken := fctx.Credentials.SecurityToken
	endpoint := os.Getenv("ROCKETMQ_ENDPOINT")
	topic := os.Getenv("TOPIC")
	instanceID := os.Getenv("INSTANCEID")
	client := mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, stsToken)
	mqProducer := client.GetProducer(instanceID, topic)
	var msg mq_http_sdk.PublishMessageRequest
	msg = mq_http_sdk.PublishMessageRequest{
		MessageBody: "hello Rocketmq",
		MessageTag:  "test greeting",
		Properties:  map[string]string{},
	}
	msg.MessageKey = "MessageKey"
	msg.Properties["a"] = "1"
	// 定时消息，定时时间为10s后，值为毫秒级别的Unix时间戳。
	msg.StartDeliverTime = time.Now().UTC().Unix()*1000 + 10*1000

	ret, err := mqProducer.PublishMessage(msg)

	if err != nil {
		return handleError(w, http.StatusBadRequest, err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	_, err = w.Write([]byte(fmt.Sprintf("Publish succ, message id : %s, messagebody md5 : %s", ret.MessageId, ret.MessageBodyMD5)))
	if err != nil {
		return handleError(w, http.StatusServiceUnavailable, err)
	}
	return nil
}

func main() {
	fc.StartHttp(HandleRequest)
}
