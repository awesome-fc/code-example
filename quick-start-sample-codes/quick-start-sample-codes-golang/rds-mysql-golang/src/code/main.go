package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
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

//connect databse
func initialize(ctx context.Context) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbEndpoint := os.Getenv("MYSQL_ENDPOINT")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")
	mySqlConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbEndpoint, dbPort, dbName)
	connection, err := sql.Open("mysql", mySqlConfig)
	db = connection
	if err != nil {
		log.Println("ERROR:Could not connect to MySql instance.", err)
	}
}

func HandleRequest(ctx context.Context) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	//insert data
	result, err := db.Exec("INSERT INTO `USERS` (`NAME`, `AGE`) VALUES (?, ?)", "wanger", 38)
	if err != nil {
		flog.Info("ERROR:", err)
		return "fail", err
	}
	newId, _ := result.LastInsertId()
	flog.Infof("New data ID added to the database is: %d \n", newId)

	//query data
	row := db.QueryRow("SELECT * FROM USERS ORDER BY ID DESC")
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		flog.Info("ERROR:", err)
		return "fail", err
	}
	flog.Info("user:", user)

	return fmt.Sprintf("user:%v", user), nil

}

//close databse
func preStop(ctx context.Context) {
	err := db.Close()
	if err != nil {
		log.Println("ERROR:Could not close MySql database.", err)
		return
	}

}
func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.RegisterPreStopFunction(preStop)
	fc.Start(HandleRequest)
}
