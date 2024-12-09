package database

import (
	"context"
	"exclusive-sample/internal/config"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase(c config.Config) (*mongo.Database, error) {
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
