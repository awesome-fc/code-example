package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	queries := req.URL.Query()
	getRowRequest := new(tablestore.GetRowRequest)
	criteria := new(tablestore.SingleRowQueryCriteria)
	putPk := new(tablestore.PrimaryKey)
	// 本示例中表格存储表名为 fc_test, 主键包含两列 region 和 id
	putPk.AddPrimaryKeyColumn("region", queries.Get("region"))
	id, err := strconv.Atoi(queries.Get("id"))
	if err != nil {
		HandleError(err)
	}
	putPk.AddPrimaryKeyColumn("id", int64(id))

	criteria.PrimaryKey = putPk
	getRowRequest.SingleRowQueryCriteria = criteria
	getRowRequest.SingleRowQueryCriteria.TableName = os.Getenv("TABLE_NAME")
	getRowRequest.SingleRowQueryCriteria.MaxVersion = 1
	getResp, err := Client.GetRow(getRowRequest)

	res := ""
	for _, col := range getResp.Columns {
		res += col.ColumnName + ": " + col.Value.(string) + ", "
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(res))
	return nil
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.StartHttp(HandleHttpRequest)
}
