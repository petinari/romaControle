package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

// var DB *mongo.Database

var Client *mongo.Client

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var e error
	Client, e = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB")))
	if e != nil {
		panic(e)
	}
}
