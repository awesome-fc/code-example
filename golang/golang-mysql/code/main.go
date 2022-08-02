package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	user User
)

type User struct {
	Id   int
	Name string
	Age  int
}

//数据库连接
func initialize(ctx context.Context) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbEndpoint := os.Getenv("MYSQL_ENDPOING")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")
	mySqlConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbEndpoint, dbPort, dbName)
	connection, err := sql.Open("mysql", mySqlConfig)
	db = connection
	if err != nil {
		log.Println("ERROR:Could not connect to MySql instance.", err)
	}
}

//数据读写操作
func HandleRequest(ctx context.Context) (string, error) {

	//插入一条数据
	result, err := db.Exec("INSERT INTO `USERS` (`NAME`, `AGE`) VALUES (?, ?)", "王二", 38)
	if err != nil {
		log.Println("ERROR:", err)
		return "fail", err
	}
	newId, _ := result.LastInsertId()
	log.Printf("New data ID added to the database is: %d \n", newId)

	//查询数据
	row := db.QueryRow("SELECT * FROM USERS ORDER BY ID DESC")
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		log.Println("ERROR:", err)
	}
	log.Println("user:", user)

	return "suc", nil

}

//数据库关闭
func preStop(ctx context.Context) {
	err := db.Close()
	if err != nil {
		log.Println("ERROR:Could not close MySql instance.", err)
		return
	}

}
func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.RegisterPreStopFunction(preStop)
	fc.Start(HandleRequest)
}
