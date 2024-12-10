package database

import (
	"context"
	"exclusive-sample/internal/config"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func ConnectDatabase(c config.Config) (*mongo.Database, error) {
	fmt.Println("name", c.DBName)
	ctx := context.TODO()
	mongoConn := options.Client().ApplyURI(c.DBUrl)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		return nil, err
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("mongo connection established")

	return mongoClient.Database(c.DBName), nil
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if mongoClient != nil {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from database: %v", err)
			return
		}
		log.Println("Successfully disconnected from MongoDB")
	}
}
