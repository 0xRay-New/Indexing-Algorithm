package database

import (
	"calculator/common"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddIndexedCollection(data common.CollectionStruct) error {
	
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"address": data.Address,
	}
	_, err := IndexedCollectionDB.UpdateOne(context.TODO(), filter, bson.M{
		"$set": bson.M{
			"address": data.Address,
			"collection_name": data.CollectionName,
			"image_url": data.ImageURL,
			"fees": data.Fees,
			"indexed_data": data.IndexedData,
		},
	}, opts)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}