package model

import (
	"market-analysis/config"
	"market-analysis/shared/mongo"
	"testing"
)

func init() {
	config.Load("../config", "config")
	mongo.InitMongo()
}
func TestInsert(t *testing.T) {

	err := (&Ticker{Name: "test"}).Insert()
	if err != nil {
		t.Log("err", err)
	}
}
