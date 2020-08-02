package core

import (
	"../db"
	"../models"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func AddCategory(data models.Category) error {
	categoryCollection := db.CategoryCollectionInit()

	createdTime := time.Now().Unix()
	id := uuid.New().String()

	data.Id = id
	data.Time = createdTime
	data.UpdatedTime = createdTime

	_, err := categoryCollection.InsertOne(db.Ctx, data)

	if err != nil {
		log.Println("Error -> ", err)
		return err
	}

	log.Println("Inserted category " + data.Name)

	return nil
}

func EditCategory(id string, data models.EditCategory) error {
	categoryCollection := db.CategoryCollectionInit()
	data.UpdatedTime = time.Now().Unix()

	filter := bson.D{
		{"_id", id},
	}

	update := bson.M{
		"$set": data,
	}

	_, err := categoryCollection.UpdateOne(db.Ctx, filter, update)

	if err != nil {
		log.Println("Error -> ", err)
		return err
	}

	log.Println("Edited category " + data.Name)
	return nil
}

func GetCategory(id string) (*models.Category, error) {
	var category models.Category
	categoryCollection := db.CategoryCollectionInit()

	filter := bson.D{
		{"_id", id},
	}

	projection := bson.D{
		{"name", 1},
	}

	findOptions := options.FindOne().SetProjection(projection)

	result := categoryCollection.FindOne(db.Ctx, filter, findOptions)

	err := result.Decode(&category)

	if err != nil {
		log.Println("Error -> ", err)
		return nil, err
	}

	return &category, nil
}

func GetCategories() ([]*models.Category, error) {
	var categories []*models.Category
	categoryCollection := db.CategoryCollectionInit()

	filter := bson.D{{}}

	projection := bson.D{
		{"name", 1},
	}

	findOptions := options.Find().SetProjection(projection)

	cur, err := categoryCollection.Find(db.Ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var category models.Category
		err := cur.Decode(&category)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

func AddItem(data models.Item) error {
	itemsCollection := db.ItemsCollectionInit()

	id := uuid.New().String()
	createdTime := time.Now().Unix()

	data.Id = id
	data.Time = createdTime
	data.UpdatedTime = createdTime

	_, err := itemsCollection.InsertOne(db.Ctx, data)

	if err != nil {
		log.Println("Error -> ", err)
		return err
	}

	log.Println("Inserted item " + data.Name)

	stock := &models.Stock{
		CreatedOn:        createdTime,
		Date:             createdTime,
		Quantity:         int64(data.Stock),
		QuantityAdjusted: 0,
		Reason:           "Opening Stock",
		ReferenceNumber:  0,
	}

	err = AddItemToStocks(id, data.Unit, *stock)

	if err != nil {
		return err
	}

	cost := &models.Cost{
		CreatedOn:            createdTime,
		Date:                 createdTime,
		CostPrice:            int64(data.CostPrice),
		AdjustedCostPrice:    0,
		SellingPrice:         int64(data.SellingPrice),
		AdjustedSellingPrice: 0,
		Mrp:                  int64(data.MRP),
		AdjustedMrp:          0,
		Reason:               "Opening Prices",
		ReferenceNumber:      0,
	}

	err = AddItemToCosts(id, *cost)

	if err != nil {
		return err
	}

	return nil
}

func EditItem(id string, data models.EditItem) error {
	itemsCollection := db.ItemsCollectionInit()

	filter := bson.D{
		{"_id", id},
	}

	update := bson.M{
		"$set": data,
	}

	_, err := itemsCollection.UpdateOne(db.Ctx, filter, update)

	if err != nil {
		log.Println("Error -> ", err)
		return err
	}

	log.Println("Edited item " + data.Name)

	return nil
}

func GetItem(id string) (*models.Item, error) {
	var item models.Item
	itemsCollection := db.ItemsCollectionInit()

	filter := bson.D{
		{"_id", id},
	}

	projection := bson.D{
		{"name", 1},
	}

	findOptions := options.FindOne().SetProjection(projection)

	result := itemsCollection.FindOne(db.Ctx, filter, findOptions)

	err := result.Decode(&item)

	if err != nil {
		log.Println("Error -> ", err)
		return nil, err
	}

	return &item, nil
}

func GetItems() ([]*models.Item, error) {
	var items []*models.Item
	itemsCollection := db.ItemsCollectionInit()

	filter := bson.D{{}}

	projection := bson.D{
		{"name", 1},
	}

	findOptions := options.Find().SetProjection(projection)

	cur, err := itemsCollection.Find(db.Ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var item models.Item
		err := cur.Decode(&item)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func AddItemToStocks(id, unit string, data models.Stock) error {
	var stock []models.Stock
	stocksCollection := db.StocksCollectionInit()

	stock = append(stock, data)

	stocks := &models.Stocks{
		Id:                id,
		AvailableQuantity: 0,
		Unit:              unit,
		LastUpdated:       data.Date,
		StocksHistory:     stock,
	}

	_, err := stocksCollection.InsertOne(db.Ctx, stocks)

	if err != nil {
		log.Println("Error while updating stock for item "+id+" -> ", err)
		return err
	}

	log.Println("Inserted stock to item " + id)

	return nil
}

func AddItemToCosts(id string, data models.Cost) error {
	var cost []models.Cost
	costsCollection := db.CostsCollectionInit()

	cost = append(cost, data)

	costs := &models.Costs{
		Id:           id,
		CostPrice:    data.Mrp,
		SellingPrice: data.SellingPrice,
		Mrp:          data.Mrp,
		LastUpdated:  data.Date,
		CostsHistory: cost,
	}

	_, err := costsCollection.InsertOne(db.Ctx, costs)

	if err != nil {
		log.Println("Error while updating Costs for item "+id+" -> ", err)
		return err
	}

	log.Println("Inserted stock to item " + id)

	return nil
}

func AdjustQuantity(id string, data models.Stock) error {
	stocksCollection := db.StocksCollectionInit()
	data.CreatedOn = time.Now().Unix()

	filter := bson.D{
		{"_id", id},
	}

	update := bson.M{
		"$set": bson.M{
			"available_quantity": data.Quantity,
			"last_updated":       data.Date,
		},
		"$push": bson.M{
			"stocks_history": data,
		},
	}

	_, err := stocksCollection.UpdateOne(db.Ctx, filter, update)

	if err != nil {
		log.Println("Error -> ", err)
		return err
	}

	log.Println("Adjusted Quantity for item " + id)

	return nil
}

func AdjustValue(id string, data models.Cost) error {
	costsCollection := db.CostsCollectionInit()
	data.CreatedOn = time.Now().Unix()

	filter := bson.D{
		{"_id", id},
	}

	update := bson.M{
		"$set": bson.M{
			"cost_price":    data.CostPrice,
			"selling_price": data.SellingPrice,
			"mrp":           data.Mrp,
			"last_updated":  data.Date,
		},
		"$push": bson.M{
			"costs_history": data,
		},
	}

	_, err := costsCollection.UpdateOne(db.Ctx, filter, update)

	if err != nil {
		log.Println("Error -> ", err)
		return err
	}

	log.Println("Adjusted Value for item " + id)

	return nil
}

func GetQuantity(id string) (*models.GetStock, error) {
	var stock models.GetStock

	stocksCollection := db.StocksCollectionInit()

	filter := bson.D{
		{"_id", id},
	}

	projection := bson.D{
		{"available_quantity", 1},
		{"unit", 1},
	}
	findOptions := options.FindOne().SetProjection(projection)

	result := stocksCollection.FindOne(db.Ctx, filter, findOptions)

	err := result.Decode(&stock)

	if err != nil {
		log.Println("Error -> ", err)
		return nil, err
	}

	return &stock, nil
}

func GetCost(id string) (*models.GetCost, error) {
	var cost models.GetCost

	costsCollection := db.CostsCollectionInit()

	filter := bson.D{
		{"_id", id},
	}

	projection := bson.D{
		{"mrp", 1},
		{"cost_price", 1},
		{"selling_price", 1},
	}
	findOptions := options.FindOne().SetProjection(projection)

	result := costsCollection.FindOne(db.Ctx, filter, findOptions)

	err := result.Decode(&cost)

	if err != nil {
		log.Println("Error -> ", err)
		return nil, err
	}

	return &cost, nil
}
