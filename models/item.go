package models

type Item struct {
	Id               string  `bson:"_id" json:"_id"`
	Time             int64   `bson:"time" json:"time"`
	UpdatedTime      int64   `bson:"updated_time" json:"updated_time"`
	CreatedBy        string  `bson:"created_by" json:"created_by"`
	Name             string  `bson:"name" json:"name"`
	Status           bool    `bson:"status" json:"status"`
	Size             float32 `bson:"size" json:"size"`
	SKU              string  `bson:"sku" json:"sku"`
	BarCode          string  `bson:"bar_code" json:"bar_code"`
	Category         string  `bson:"category" json:"category"`
	Unit             string  `bson:"unit" json:"unit"`
	SalesInformation bool    `bson:"sales_information" json:"sales_information"`
	MRP              float32 `bson:"mrp,truncate" json:"mrp"`
	CostPrice        float32 `bson:"cost_price,truncate" json:"cost_price"`
	SellingPrice     float32 `bson:"selling_price,truncate" json:"selling_price"`
	TrackInventory   bool    `bson:"track_inventory" json:"track_inventory"`
	Stock            float32 `bson:"stock,truncate" json:"stock"`
	ReorderPoint     float32 `bson:"reorder_point" json:"reorder_point"`
	PreferredVendor  string  `bson:"preferred_vendor" json:"preferred_vendor"`
}

type EditItem struct {
	UpdatedTime     int64   `bson:"updated_time" json:"updated_time"`
	Name            string  `bson:"name" json:"name"`
	Status          bool    `bson:"status" json:"status"`
	Size            float32 `bson:"size" json:"size"`
	SKU             string  `bson:"sku" json:"sku"`
	BarCode         string  `bson:"bar_code" json:"bar_code"`
	Category        string  `bson:"category" json:"category"`
	Unit            string  `bson:"unit" json:"unit"`
	ReorderPoint    float32 `bson:"reorder_point" json:"reorder_point"`
	PreferredVendor string  `bson:"preferred_vendor" json:"preferred_vendor"`
}
