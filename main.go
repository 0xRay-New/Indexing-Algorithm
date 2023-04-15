package main

import (
	"bytes"
	"calculator/calculator"
	"calculator/common"
	"calculator/database"
	"calculator/webhooks"
	"encoding/json"
	"os"

	"github.com/andybalholm/brotli"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// inits database
	err := database.InitMongoDb()
	
	if err != nil {
		panic(err)
	}

	// inits server
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// indexes a given collection
	// example url: http://localhost:5000/api/v1/index&collection=<collection slug here>
	app.Get("/api/v1/index", func(c *fiber.Ctx) error {
		collectionName := c.Query("collection")

		data := calculator.GetCollectionData(collectionName)

		// checks if collection has already been indexed
		if common.CheckIfCollectionInIndexedCollections(collectionName) {
			return c.Status(400).JSON(fiber.Map{
				"success": true,
				"message": "collection is indexed",
			})
		}

		// gets collection metadata
		err := data.GetMetadata()
		if err != nil {
			if err.Error() == "collection not found" {
				return c.Status(404).JSON(fiber.Map{
					"success": false,
					"message": "collection not found",
				})
			} else {
				return c.Status(400).JSON(fiber.Map{
					"success": false,
					"message": err.Error(),
				})
			}
		}

		// indexes NFTs by metadata
		data.CheckFirstIndex()
		data.GetAssetsRequest()
		data.SortNFTS()
		for i := range data.SortedNFTScores {
			data.NFTScores[data.SortedNFTScores[i].TokenID]["r"] = i+1
		}
		common.IndexedCollections = append(common.IndexedCollections, collectionName)

		// brotli encodes collection to decrease database storage usage
		d, _ := json.Marshal(data.NFTScores)
		var b bytes.Buffer
		w := brotli.NewWriter(&b)
		w.Write(d)
		w.Close()

		newData := common.CollectionStruct {
			Address:        data.Address,
			CollectionName: collectionName,
			ImageURL:       data.ImageURL,
			Fees:           float64(data.Fees)/float64(10000),
			IndexedData:    b.Bytes(),
			Count: 		data.Count,
			IndexCount: 	len(data.NFTScores),
		}
		
		database.AddIndexedCollection(newData)

		webhooks.SendIndexedCollectionWebhook(newData)
		return c.Status(200).JSON(fiber.Map{
			"address": data.Address,
			"collection_name": collectionName,
			"image_url": data.ImageURL,
			"fees": float64(data.Fees)/float64(10000),
			"success": true,
			"message": "Collection indexed successfully",
		})
	})

	app.Get("/api/v1/reindex", func(c *fiber.Ctx) error {
		address := c.Query("address")
		collectionName := c.Query("collection")
		if address == "" {
			data := calculator.GetCollectionData(collectionName)
			err := data.GetMetadata()
			if err != nil {
				if err.Error() == "collection not found" {
					return c.Status(404).JSON(fiber.Map{
						"success": false,
						"message": "collection not found",
					})
				} else {
					return c.Status(400).JSON(fiber.Map{
						"success": false,
						"message": err.Error(),
					})
				}
			}
			data.CheckFirstIndex()
			data.GetAssetsRequest()
			data.SortNFTS()
			for i := range data.SortedNFTScores {
				data.NFTScores[data.SortedNFTScores[i].TokenID]["r"] = i+1
			}
			common.IndexedCollections = append(common.IndexedCollections, collectionName)
			d, _ := json.Marshal(data.NFTScores)
			var b bytes.Buffer
			w := brotli.NewWriter(&b)
			w.Write(d)
			w.Close()

			newData := common.CollectionStruct {
				Address:        data.Address,
				CollectionName: collectionName,
				ImageURL:       data.ImageURL,
				Fees:           float64(data.Fees)/float64(10000),
				IndexedData:    b.Bytes(),
				Count: 		data.Count,
				IndexCount: 	len(data.NFTScores),
			}
			
			database.AddIndexedCollection(newData)

			webhooks.SendIndexedCollectionWebhook(newData)
			return c.Status(200).JSON(fiber.Map{
				"address": data.Address,
				"collection_name": collectionName,
				"image_url": data.ImageURL,
				"fees": float64(data.Fees)/float64(10000),
				"success": true,
				"message": "Collection indexed successfully",
			})
		}
		return c.SendStatus(400)
	})

	

	app.Listen(":" + port)
}