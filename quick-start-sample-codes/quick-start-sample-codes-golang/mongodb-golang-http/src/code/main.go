package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

var Client *mongo.Client

func initialize(ctx context.Context) {
	url := os.Getenv("MONGO_URL")

	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		HandleError(err)
	}
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	queries := req.URL.Query()
	database := os.Getenv("MONGO_DATABASE")
	coll := Client.Database(database).Collection("users")

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: queries.Get("name")}}).Decode(&result)
	if err != nil {
		HandleError(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	json_res, err := json.Marshal(result)
	if err != nil {
		HandleError(err)
	}
	w.Write(json_res)
	return nil
}

func preStop(ctx context.Context) {
	if err := Client.Disconnect(context.TODO()); err != nil {
		HandleError(err)
	}
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.RegisterPreStopFunction(preStop)
	fc.StartHttp(HandleHttpRequest)
}
