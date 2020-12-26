package db

import (
	"context"
	"tara/api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Cfg.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	Client = client

	return client, err
}
