package models

type Category struct {
	Id           string `bson:"_id" json:"_id"`
	Time         int64  `bson:"time" json:"time"`
	Name         string `bson:"name" json:"name"`
	Description  string `bson:"description" json:"description"`
	Parent       string `bson:"parent" json:"parent"`
	Unit         string `bson:"unit" json:"unit"`
	Manufacturer string `bson:"manufacture" json:"manufacture"`
	Brand        string `bson:"brand" json:"brand"`
}
