package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDB() (mongoClient *mongo.Client, err error) {
	ctx := context.TODO()
	URI := "mongodb://localhost:27017"
	connectionOpts := options.Client().ApplyURI(URI)

	mongoClient, err = mongo.Connect(ctx, connectionOpts)
	if err != nil {
		return nil, err
	}

	connectionOpts.SetMaxPoolSize(10)
	connectionOpts.SetMinPoolSize(5)
	connectionOpts.SetMaxConnIdleTime(10 * time.Second)

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return mongoClient, err
}
