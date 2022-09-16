package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	user User
)

type User struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handleError(w http.ResponseWriter, err error) error {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(err.Error()))
	return err
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

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(w, err)
	}
	var bodyJson User
	err = json.Unmarshal(body, &bodyJson)
	if err != nil {
		handleError(w, err)
	}
	//insert data
	result, err := db.Exec("INSERT INTO `USERS` (`NAME`, `AGE`) VALUES (?, ?)", bodyJson.Name, bodyJson.Age)
	if err != nil {
		handleError(w, err)
	}
	newId, _ := result.LastInsertId()
	log.Printf("New data ID added to the database is: %d \n", newId)

	//query data
	row := db.QueryRow("SELECT * FROM USERS ORDER BY ID DESC")
	err = row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	log.Println("user:", user)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	_, err = w.Write([]byte("succ"))
	if err != nil {
		handleError(w, err)
	}
	return nil

}

//close databse
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
	fc.StartHttp(HandleHttpRequest)
}
