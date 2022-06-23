package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func HandleRequest(ctx context.Context, event OssEvent) (string, error) {
	// 获取密钥信息，执行前，确保函数所在的服务配置了角色信息，并且角色需要拥有 AliyunOSSFullAccess权限
	fctx, _ := fccontext.FromContext(ctx)
	creds := fctx.Credentials
	// endpoint 填写 Bucket 对应的 Endpoint，以华南1（深圳）为例，填写为 https://oss-cn-shenzhen-internal.aliyuncs.com。其它 Region 请按实际情况填写。
	endPoint := "https://oss-cn-shenzhen-internal.aliyuncs.com"
	// 创建 OssClient 实例,为了防止暴露用户 AccessKeyId、AccessKeySecret 字段，这里使用 STS 新建一个实例
	ossClient, err := oss.New(endPoint, creds.AccessKeyId, creds.AccessKeySecret, oss.SecurityToken(creds.SecurityToken))
	if err != nil {
		HandleError(err)
	}
	//获取 oss 的 Bucket
	buckName := event.Events[0].Oss.Bucket.Name
	bucket, err := ossClient.Bucket(buckName)
	if err != nil {
		HandleError(err)
	}
	//获得上传至 oss 图片的名称
	sourceImageName := event.Events[0].Oss.Object.Key
	//判断这个图片是否已经被压缩过，如果已经是压缩图片则跳过，防止循环压缩
	if strings.HasPrefix(sourceImageName, "w_100_h_100") {
		return "skip", nil
	}
	// 指定压缩处理后的图片名称。
	targetImageName := "w_100_h_100_" + event.Events[0].Oss.Object.Key
	//将图片缩放为固定宽高100 px后保存在oss中
	style := "image/resize,m_fixed,w_100,h_100"
	process := fmt.Sprintf("%s|sys/saveas,o_%v,b_%v", style, base64.URLEncoding.EncodeToString([]byte(targetImageName)), base64.URLEncoding.EncodeToString([]byte(buckName)))
	_, err = bucket.ProcessObject(sourceImageName, process)
	if err != nil {
		HandleError(err)
	}
	return "done", nil
}

func main() {
	fc.Start(HandleRequest)
}
