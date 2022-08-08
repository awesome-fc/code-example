package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

var Client *tablestore.TableStoreClient

func initialize(ctx context.Context) {
	fctx, _ := fccontext.FromContext(ctx)
	var (
		endPoint        string = os.Getenv("ENDPOINT")
		instanceName    string = os.Getenv("INSTANCE_NAME")
		accessKey       string = fctx.Credentials.AccessKeyId
		accessKeySecret string = fctx.Credentials.AccessKeySecret
		stsToken        string = fctx.Credentials.SecurityToken
	)
	Client = tablestore.NewClientWithConfig(endPoint, instanceName, accessKey, accessKeySecret, stsToken, nil)
}

func HandleRequest(ctx context.Context) (*tablestore.GetRowResponse, error) {
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	// 本示例中表格存储表名为 fc_test, 主键包含两列 region 和 id
	putPk.AddPrimaryKeyColumn("region", "abc")
	putPk.AddPrimaryKeyColumn("id", int64(1))

	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = os.Getenv("TABLE_NAME")
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
	getResp, err := Client.GetRow(getRowRequest)

	if err != nil {
		HandleError(err)
	}
	return getResp, nil
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.Start(HandleRequest)
}
