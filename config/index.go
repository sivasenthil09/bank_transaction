package config

import (
	"context"
	"log"
	"time"

	"github.com/sivasenthil09/bank_transaction/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDatabase() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoconnection := options.Client().ApplyURI(constants.ConnectionString)

	MongoClient, err := mongo.Connect(ctx, mongoconnection)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	

	return MongoClient, nil
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
