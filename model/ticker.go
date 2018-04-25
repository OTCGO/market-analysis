package model

import (
	"market-analysis/shared/mongo"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
{
	"id": "bitcoin",
	"name": "Bitcoin",
	"symbol": "BTC",
	"rank": "1",
	"price_usd": "8848.78",
	"price_btc": "1.0",
	"24h_volume_usd": "6603630000.0",
	"market_cap_usd": "150374061310",
	"available_supply": "16993762.0",
	"total_supply": "16993762.0",
	"max_supply": "21000000.0",
	"percent_change_1h": "0.33",
	"percent_change_24h": "0.09",
	"percent_change_7d": "7.74",
	"last_updated": "1524450272",
	"price_cny": "55763.241804",
	"24h_volume_cny": "41614755534.0",
	"market_cap_cny": "947627259566"
}
*/
type Ticker struct {
	Id_              bson.ObjectId `bson:"_id"`
	Name             string        `json:"name" bson:"name"`
	Symbol           string        `json:"symbol" bson:"symbol"`
	Rank             string        `json:"rank" bson:"rank"`
	PriceUsd         string        `json:"price_usd" bson:"price_usd"`
	PriceBtc         string        `json:"price_btc" bson:"price_btc"`
	VolumeUsd24h     string        `json:"24h_volume_usd" bson:"24h_volume_usd"`
	MarketCapUsd     string        `json:"market_cap_usd" bson:"market_cap_usd"`
	AvailableSupply  string        `json:"available_supply" bson:"available_supply"`
	TotalSupply      string        `json:"total_supply" bson:"total_supply"`
	MaxSupply        string        `json:"max_supply" bson:"max_supply"`
	PercentChange1h  string        `json:"percent_change_1h" bson:"percent_change_1h"`
	PercentChange24h string        `json:"percent_change_24h" bson:"percent_change_24h"`
	PercentChange7h  string        `json:"percent_change_7d" bson:"percent_change_7d"`
	PriceCny         string        `json:"price_cny" bson:"price_cny"`
	VolumeCny24h     string        `json:"24h_volume_cny" bson:"24h_volume_cny"`
	MarketCapCny     string        `json:"market_cap_cny" bson:"market_cap_cny"`
	LastUpdated      string        `json:"last_updated" bson:"last_updated"`
}

func (t *Ticker) CreateIndex() (err error) {
	session := mongo.GetSession()
	// close session
	defer session.Clone()
	err = session.DB(mongo.DataBase).C("Ticker").EnsureIndex(mgo.Index{
		Key:    []string{"last_updated", "name"},
		Unique: true,
	})
	err = session.DB(mongo.DataBase).C("Ticker").EnsureIndex(mgo.Index{
		Key: []string{"rank"},
	})
	err = session.DB(mongo.DataBase).C("Ticker").EnsureIndex(mgo.Index{
		Key: []string{"name"},
	})
	return err
}

func (t *Ticker) Insert() (err error) {
	// fmt.Println("t", t.Id_)
	if t.Id_ == "" {
		t.Id_ = bson.NewObjectId()
	}

	session := mongo.GetSession()
	// close session
	defer session.Clone()
	err = session.DB(mongo.DataBase).C("Ticker").Insert(t)
	return err
}
