package sendMessage

import (
	"fmt"
	"strconv"
	"time"

	mq_http_sdk "github.com/aliyunmq/mq-http-go-sdk"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	// 设置HTTP接入域名。
	endpoint := "http://143xxxxxxxxxxxx.mqrest.cn-shanghai.aliyuncs.com"
	// AccessKey ID，阿里云身份验证。获取方式，请参见本文前提条件中的获取AccessKey。
	accessKey := ""
	// AccessKey Secret，阿里云身份验证。获取方式，请参见本文前提条件中的获取AccessKey。
	secretKey := ""
	// 所属的Topic。
	topic := ""
	// Topic所属实例ID，默认实例为空。
	instanceId := ""

	client := mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, "")

	mqProducer := client.GetProducer(instanceId, topic)
	// 循环发送4条消息。
	for i := 0; i < 2; i++ {
		var msg mq_http_sdk.PublishMessageRequest
		if i%2 == 0 {
			msg = mq_http_sdk.PublishMessageRequest{
				MessageBody: "hello mq!",         //消息内容。
				MessageTag:  "greeting",          // 消息标签。
				Properties:  map[string]string{}, // 消息属性。
			}
			// 设置Key。
			msg.MessageKey = "MessageKey"
			// 设置属性。
			msg.Properties["a"] = strconv.Itoa(i)
		} else {
			msg = mq_http_sdk.PublishMessageRequest{
				MessageBody: "i am sending messages!", //消息内容。
				MessageTag:  "demo",                   // 消息标签。
				Properties:  map[string]string{},      // 消息属性。
			}
			// 设置属性。
			msg.Properties["a"] = strconv.Itoa(i)
			// 定时消息，定时时间为10s后，值为毫秒级别的Unix时间戳。
			msg.StartDeliverTime = time.Now().UTC().Unix()*1000 + 10*1000
		}
		ret, err := mqProducer.PublishMessage(msg)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Publish ---->\n\tMessageId:%s, BodyMD5:%s, \n", ret.MessageId, ret.MessageBodyMD5)
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	fmt.Printf("all messages sent!")
}
