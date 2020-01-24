package dependencies

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.MONGO_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	Client = client

	return client, err
}
