package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitializeMongo(uri, username, password string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cred := options.Credential{
		Username: username,
		Password: password,
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(cred))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
