package models

import (
	"context"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()
var RDB = InitRedis()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// mongoDB若有密码请Client()调用SetAuth()方法
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Println("Connection MongoDB Error: ", err)
		return nil
	}
	return client.Database("im")
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}
