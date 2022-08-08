package main

import (
	"context"
	"fmt"
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

func HandleRequest(ctx context.Context) (bson.M, error) {
	database := os.Getenv("MONGO_DATABASE")
	coll := Client.Database(database).Collection("users")

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"name", "张三"}}).Decode(&result)
	if err != nil {
		HandleError(err)
	}

	return result, nil
}

func preStop(ctx context.Context) {
	if err := Client.Disconnect(context.TODO()); err != nil {
		HandleError(err)
	}
}

func main() {
	fc.RegisterInitializerFunction(initialize)
	fc.RegisterPreStopFunction(preStop)
	fc.Start(HandleRequest)
}
