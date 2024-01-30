package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	Db           *mongo.Database
	Client       *mongo.Client
	MdCollection *mongo.Collection
}

func NewDatabase(ctx context.Context) (*Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, fmt.Errorf("mongo client creation: %w", err)
	}
	db := client.Database("profile")
	mdCollection := db.Collection("md")

	return &Database{
		Db:           db,
		Client:       client,
		MdCollection: mdCollection,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.Ping(ctx, readpref.Primary())
}

func (d *Database) Close(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}
