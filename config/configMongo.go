package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// var DB *mongo.Database

var Client *mongo.Client

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//db, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	//DB = db.Database("bancoUnico")
	var e error
	Client, e = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if e != nil {
		panic(e)
	}
}
