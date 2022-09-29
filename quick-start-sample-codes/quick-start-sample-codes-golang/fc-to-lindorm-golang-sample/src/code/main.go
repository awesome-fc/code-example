// 本代码样例主要实现以下功能:
// 初始化 lindorm 连接
// 从环境变量中获取 lindorm 宽表引擎 table 名字
// 创建此表并向表中插入一条数据
// 读表中所有数据，查看是否符合预期
//
// This sample code is mainly doing the following things:
// Initialize lindorm connection
// Get the lindorm table name from environment variables
// Create this table and insert data into the table
// Read all the data in the table to see if it is as expected
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	avatica "github.com/apache/calcite-avatica-go/v5"
)

var db *sql.DB

func initialize(ctx context.Context) {
	fmt.Println("this is initialize handler")
	lindormUserName := os.Getenv("LindormUserName")
	lindormPassword := os.Getenv("LindormPassword")
	databaseURL := os.Getenv("DatabaseURL")

	conn := avatica.NewConnector(databaseURL).(*avatica.Connector)
	conn.Info = map[string]string{
		"user":     lindormUserName, // 数据库用户名
		"password": lindormPassword, // 数据库密码
		"database": "default",       // 初始化连接指定的默认database
	}
	db = sql.OpenDB(conn)
}

func HandleRequest(ctx context.Context) (string, error) {
	if db == nil {
		return "", fmt.Errorf("lindorm connection not initialized.")
	}
	sqlTableName := os.Getenv("SQLTableName")

	// 如果输入的 table 存在，则先删除，保证表数据清洁
	// If the input table exists, delete it first to ensure that the table data is clean
	_, err := db.Exec(fmt.Sprintf("drop table if exists %s", sqlTableName))
	if err != nil {
		return "", err
	}

	// 创建新表，新表共有两列：c1、c2
	// Create a new table, the new table has two columns: c1, c2
	_, err = db.Exec(fmt.Sprintf("create table if not exists %s(c1 int, c2 int, primary key(c1))", sqlTableName))
	if err != nil {
		return "", err
	}

	// 如果表创建成功，则需在请求执行结束后删除。
	// If the table is created successfully, it needs to be deleted after the request execution ends.
	defer deleteTable(sqlTableName)

	// 向表中插入数据，两列数据分别为 20 和 30
	// Insert data into the table, the two columns of data are 20 and 30
	_, err = db.Exec(fmt.Sprintf("upsert into %s(c1,c2) values(20,30)", sqlTableName))
	if err != nil {
		return "", err
	}

	// 查询表中的数据并输出
	// Query the data in the table and output
	querySql := fmt.Sprintf("select * from %s", sqlTableName)
	rows, err := db.Query(querySql)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var c1 int
	var c2 string
	for rows.Next() {
		err = rows.Scan(&c1, &c2)
		if err != nil {
			return "", err
		}
		fmt.Printf("c1: %d, c2: %s\n", c1, c2)
	}
	return "success", nil
}

// 删除表
// delete table
func deleteTable(sqlTableName string) {
	_, err := db.Exec(fmt.Sprintf("OFFLINE TABLE %s", sqlTableName))
	if err != nil {
		fmt.Println(fmt.Sprintf("Warning: offline table failed due to %v", err))
		return
	}

	_, err = db.Exec(fmt.Sprintf("drop table if exists %s", sqlTableName))
	if err != nil {
		fmt.Println(fmt.Sprintf("Warning: drop table failed due to %v", err))
	}
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.Start(HandleRequest)
}
