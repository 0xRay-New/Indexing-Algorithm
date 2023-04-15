package database

import (
	"calculator/common"
	"context"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var (
	MongoDbName = "Users-IndexedData"
	MongoIndexedCollectionName = "Indexed-Data"
)

var (
	IndexedCollectionDB *mongo.Collection
)

var (
	Client *mongo.Client
)

func InitMongoDb() error {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(common.MongoDbConnectionURL).SetServerAPIOptions(serverAPIOptions)
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return err
	}
	Client = MongoClient
	log.Println("started db")

	IndexedCollectionDB = Client.Database(MongoDbName).Collection(MongoIndexedCollectionName)

	return nil
}