package models

type Timer struct {
	Dsmesid  string `bson:"dsmesid"`
	Dschatid string `bson:"dschatid"`
	Tgmesid  string `bson:"tgmesid"`
	Tgchatid string `bson:"tgchatid"`
	Timed    int    `bson:"timed"`
}
