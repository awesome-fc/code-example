package main

import (
	"context"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/fxamacker/cbor/v2"
	"log"
)

/*
* 本代码样例主要实现以下功能:
** 打印 event


* This sample code is mainly doing the following things:
** Print event
 */

type OTSEvent struct {
	Version *string           `json:"Version"`
	Records []*OTSEventRecord `json:"Records"`
}

type OTSEventRecord struct {
	Type       *string       `json:"Type"`
	Info       *Info         `json:"Info"`
	PrimaryKey []*PrimaryKey `json:"PrimaryKey"`
	Columns    []*Column     `json:"Columns"`
}

type Info struct {
	Timestamp *int `json:"Timestamp"`
}

type PrimaryKey struct {
	ColumnName *string      `json:"ColumnName"`
	Value      *interface{} `json:"Value"`
}

type Column struct {
	Type       *string      `json:"Type"`
	ColumnName *string      `json:"ColumnName"`
	Value      *interface{} `json:"Value"`
	Timestamp  *int         `json:"Timestamp"`
}

func main() {
	fc.Start(HandleRequest)
}

// event 示例见文档：https://help.aliyun.com/document_detail/169672.html，event结构如下所示：
// {
//     "Version": "Sync-v1",
//     "Records": [
//         {
//             "Type": "PutRow",
//             "Info": {
//                 "Timestamp": 1506416585740836
//             },
//             "PrimaryKey": [
//                 {
//                     "ColumnName": "pk_0",
//                     "Value": 1506416585881590900
//                 },
//                 {
//                     "ColumnName": "pk_1",
//                     "Value": "2017-09-26 17:03:05.8815909 +0800 CST"
//                 },
//                 {
//                     "ColumnName": "pk_2",
//                     "Value": 1506416585741000
//                 }
//             ],
//             "Columns": [
//                 {
//                     "Type": "Put",
//                     "ColumnName": "attr_0",
//                     "Value": "hello_table_store",
//                     "Timestamp": 1506416585741
//                 },
//                 {
//                     "Type": "Put",
//                     "ColumnName": "attr_1",
//                     "Value": 1506416585881590900,
//                     "Timestamp": 1506416585741
//                 }
//             ]
//         }
//     ]
// }

func HandleRequest(ctx context.Context, eventBytes []byte) (string, error) {
	var event OTSEvent
	err := cbor.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Println("error!", err)
	}
	records := event.Records
	version := event.Version
	fmt.Printf("Version: %s\n", *version)
	for _, v := range records {
		record := *v
		fmt.Printf("Type: %s\n", *record.Type)
		fmt.Printf("Timestamp: %d\n", *record.Info.Timestamp)
		for _, pk := range record.PrimaryKey {
			fmt.Printf("pk name: %s\n", *pk.ColumnName)
		}
		for _, col := range record.Columns {
			fmt.Printf("column type: %s\n", *col.Type)
			fmt.Printf("column name: %s\n", *col.ColumnName)
		}
	}
	return "ok", nil
}
