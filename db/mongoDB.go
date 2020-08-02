package db

import (
	configuration "../config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var config = configuration.Init()
var client *mongo.Client = nil
var Ctx = context.TODO()

var categoryCollection *mongo.Collection = nil
var itemsCollection *mongo.Collection = nil
var stocksCollection *mongo.Collection = nil
var costsCollection *mongo.Collection = nil

func Init() {
	if client == nil {
		clientOptions := options.Client().ApplyURI(config.DB_Url)
		var err error = nil
		client, err = mongo.Connect(Ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Connected to MongoDB.")
		}
	}
}

func CategoryCollectionInit() *mongo.Collection {
	if categoryCollection == nil {
		categoryCollection = client.Database(config.DB_Name).Collection("category")
	}
	return categoryCollection
}

func ItemsCollectionInit() *mongo.Collection {
	if itemsCollection == nil {
		itemsCollection = client.Database(config.DB_Name).Collection("items")
	}
	return itemsCollection
}

func StocksCollectionInit() *mongo.Collection {
	if stocksCollection == nil {
		stocksCollection = client.Database(config.DB_Name).Collection("stocks")
	}
	return stocksCollection
}

func CostsCollectionInit() *mongo.Collection {
	if costsCollection == nil {
		costsCollection = client.Database(config.DB_Name).Collection("costs")
	}
	return costsCollection
}
