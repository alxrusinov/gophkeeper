package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo - repository with MongoDB
type Mongo struct {
	client *mongo.Client
}

// Disconnect is a method closing mongo db connection
func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

// NewMongo - create instance of MongoDB repository
func NewMongo(ctx context.Context, address string) (*Mongo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return &Mongo{
		client: client,
	}, nil
}
