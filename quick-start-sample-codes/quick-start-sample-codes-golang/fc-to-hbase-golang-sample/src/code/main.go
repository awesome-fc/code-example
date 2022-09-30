// 本代码样例主要实现以下功能:
// 初始化 Hbase client
// 从环境变量中获取 Hbase table 名字
// 创建表格并向表中插入一条数据
// 读表中插入行的数据
// 清理该行数据
// 删除表格
//
// This sample code is mainly doing the following things:
// Initialize Hbase client
// Get the Hbase table name from environment variables
// Create this table and insert data into the table
// Read the data of the row in the table
// Clear the data inserted before
// delete the table
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/sdming/goh"
	"github.com/sdming/goh/Hbase"
)

var hbaseClient *goh.HClient

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.Start(HandleRequest)
}

func initialize(ctx context.Context) {
	fmt.Println("this is initialize handler")
	hbaseThriftURL := os.Getenv("HBaseThriftURL")
	hbaseConnPort := os.Getenv("HBaseConnPort")
	address := fmt.Sprintf("%s:%s", hbaseThriftURL, hbaseConnPort)
	fmt.Println(address)

	tcpClinet, err := goh.NewTcpClient(address, goh.TBinaryProtocol, false)
	if err != nil {
		fmt.Println(err)
	}
	hbaseClient = tcpClinet
}

func HandleRequest(ctx context.Context) (string, error) {
	if hbaseClient == nil {
		return "", fmt.Errorf("hbase client not initialized.")
	}

	if err := hbaseClient.Open(); err != nil {
		fmt.Println(err)
		return "", err
	}

	defer hbaseClient.Close()

	tableName := os.Getenv("TableName")

	columns := make([]string, 1)
	columns[0] = "info:c1"

	// 根据给出的 Table 名字试图创建 Table
	fmt.Printf("CreateTable: %s", tableName)
	cols := make([]*goh.ColumnDescriptor, 1)
	cols[0] = goh.NewColumnDescriptorDefault("info:c1")
	if exist, err := hbaseClient.CreateTable(tableName, cols); err != nil {
		return "", err
	} else {
		fmt.Println(exist)
	}

	// 列出当前现存的 Table 列表
	fmt.Print("GetTableNames:")
	if data, err := hbaseClient.GetTableNames(); err != nil {
		fmt.Println(err)
	} else {
		dump(data)
	}

	// 确认当前 Table 是否为 Enable 状态
	fmt.Print("IsTableEnabled:")
	tableEnabled, err := hbaseClient.IsTableEnabled(tableName)
	if err != nil {
		return "", err
	}
	fmt.Println(tableEnabled)

	// 如果 Table 状态不为 Enable，则将其变更为 Enable
	if !tableEnabled {
		fmt.Print("EnableTable:")
		if err := hbaseClient.EnableTable(tableName); err != nil {
			return "", err
		}
	}

	// 向 Table 中插入一行
	mutations := make([]*Hbase.Mutation, 1)
	mutations[0] = goh.NewMutation(columns[0], []byte("value3-mutation"))

	fmt.Print("MutateRow:")
	fmt.Println(hbaseClient.MutateRow(tableName, []byte("row1"), mutations, nil))

	// 获取插入行的数据
	fmt.Print("GetRowWithColumns:")
	if data, err := hbaseClient.GetRowWithColumns(tableName, []byte("row1"), columns, nil); err != nil {
		return "", err
	} else {
		printRows(data)
	}

	// 清除插入行数据
	fmt.Print("DeleteAllRow:")
	fmt.Println(hbaseClient.DeleteAllRow(tableName, []byte("row1"), nil))

	// 将 Table 状态变更为 Disable
	fmt.Print("DisableTable:")
	fmt.Println(hbaseClient.DisableTable(tableName))

	// 删除该 Table
	fmt.Print("DeleteTable:")
	fmt.Println(hbaseClient.DeleteTable(tableName))

	return "done", nil
}

func dump(data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}
	fmt.Println(string(b))
}

func printRows(data []*Hbase.TRowResult) {
	if data == nil {
		fmt.Println("<nil>")
	}

	l := len(data)
	fmt.Println("[]*Hbase.TRowResult len:", l)
	for i, x := range data {
		fmt.Println(i, string(x.Row), "\n[")
		for k, v := range x.Columns {
			fmt.Println("\t", k, ":", string(v.Value), v.Timestamp)
		}
		fmt.Println("]")
	}
}
