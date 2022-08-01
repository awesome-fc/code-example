package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

/*
* 本代码样例主要实现以下功能:
** 从 event 中解析出 OSS 事件触发相关信息
** 根据以上获取的信息，初始化 OSS bucket 客户端
** 从 OSS bucket 下载即将被处理的目标图片
** 改变目标图片尺寸
** 将处理过的图片上传到 OSS bucket 下的 processed 目录


* This sample code is mainly doing the following things:
** Get OSS processing related information from event
** Initiate OSS client with target bucket
** Download the target image to be processed from bucket
** Resize the image
** Upload the processed image copy into the same bucket's processed folder
 */

func main() {
	fc.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event events.OssEvent) (string, error) {
	// 获取密钥信息，执行前，确保函数所在的服务配置了角色信息，并且角色需要拥有 AliyunOSSFullAccess权限
	fctx, _ := fccontext.FromContext(ctx)
	creds := fctx.Credentials
	fmt.Printf("The creds passed into function are %+v \n", creds)
	fmt.Printf("The event passed into function is %+v \n", event)
	ossInfo := event.Events[0].Oss
	region := *event.Events[0].Region
	// endpoint 生成 Bucket 对应的 Endpoint，以华南1（深圳）为例，endpoint 为 https://oss-cn-shenzhen-internal.aliyuncs.com。
	endPoint := "https://oss-" + region + "-internal.aliyuncs.com"
	// 创建 OssClient 实例,为了防止暴露用户 AccessKeyId、AccessKeySecret 字段，这里使用 STS 新建一个实例
	ossClient, err := oss.New(endPoint, creds.AccessKeyId, creds.AccessKeySecret, oss.SecurityToken(creds.SecurityToken))
	if err != nil {
		cErr := fmt.Errorf("see error when setting up oss client %+v", err)
		fmt.Printf("see error when setting up oss client %+v", err)
		return "", cErr
	}
	// 获取 oss 的 Bucket
	buckName := *ossInfo.Bucket.Name
	bucket, err := ossClient.Bucket(buckName)
	if err != nil {
		bErr := fmt.Errorf("see error when setting up oss client %+v", err)
		fmt.Printf("see error when setting up oss bucket %+v", err)
		return "", bErr
	}
	//获得上传至 oss 图片的名称
	sourceImageName := *ossInfo.Object.Key
	// 指定压缩处理后的图片路径，将其放入 processed 文件夹中。
	targetImageName := "processed/" + sourceImageName
	// 将图片缩放为固定宽高100 px后保存在oss中
	style := "image/resize,m_fixed,w_100,h_100"
	process := fmt.Sprintf("%s|sys/saveas,o_%v,b_%v", style, base64.URLEncoding.EncodeToString([]byte(targetImageName)), base64.URLEncoding.EncodeToString([]byte(buckName)))
	_, err = bucket.ProcessObject(sourceImageName, process)
	if err != nil {
		pErr := fmt.Errorf("see error when setting up oss client %+v", err)
		fmt.Printf("see error when processing image %+v", err)
		return "", pErr
	}
	return "done", nil
}
