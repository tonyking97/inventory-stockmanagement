package models

type Stocks struct {
	Id                string  `bson:"_id" json:"_id"`
	AvailableQuantity int64   `bson:"available_quantity" json:"available_quantity"`
	Unit              string  `bson:"unit" json:"unit"`
	LastUpdated       int64   `bson:"last_updated" json:"last_updated"`
	StocksHistory     []Stock `bson:"stocks_history" json:"stocks_history"`
}

type Stock struct {
	CreatedOn        int64  `bson:"created_on" json:"created_on"`
	Date             int64  `bson:"date" json:"date"`
	Quantity         int64  `bson:"quantity" json:"quantity"`
	QuantityAdjusted int64  `bson:"quantity_adjusted" json:"quantity_adjusted"`
	Reason           string `bson:"reason" json:"reason"`
	ReferenceNumber  int64  `bson:"reference_number" json:"reference_number"`
}

type GetStock struct {
	AvailableQuantity int64  `bson:"available_quantity" json:"available_quantity"`
	Unit              string `bson:"unit" json:"unit"`
}
