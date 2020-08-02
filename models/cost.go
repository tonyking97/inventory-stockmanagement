package models

type Costs struct {
	Id           string `bson:"_id" json:"_id"`
	CostPrice    int64  `bson:"cost_price" json:"cost_price"`
	SellingPrice int64  `bson:"selling_price" json:"selling_price"`
	Mrp          int64  `bson:"mrp" json:"mrp"`
	LastUpdated  int64  `bson:"last_updated" json:"last_updated"`
	CostsHistory []Cost `bson:"costs_history" json:"costs_history"`
}

type Cost struct {
	CreatedOn            int64  `bson:"created_on" json:"created_on"`
	Date                 int64  `bson:"date" json:"date"`
	CostPrice            int64  `bson:"cost_price" json:"cost_price"`
	AdjustedCostPrice    int64  `bson:"adjusted_cost_price" json:"adjusted_cost_price"`
	SellingPrice         int64  `bson:"selling_price" json:"selling_price"`
	AdjustedSellingPrice int64  `bson:"adjusted_selling_price" json:"adjusted_selling_price"`
	Mrp                  int64  `bson:"mrp" json:"mrp"`
	AdjustedMrp          int64  `bson:"adjusted_mrp" json:"adjusted_mrp"`
	Reason               string `bson:"reason" json:"reason"`
	ReferenceNumber      int64  `bson:"reference_number" json:"reference_number"`
}

type GetCost struct {
	CostPrice    int64 `bson:"cost_price" json:"cost_price"`
	SellingPrice int64 `bson:"selling_price" json:"selling_price"`
	Mrp          int64 `bson:"mrp" json:"mrp"`
}
