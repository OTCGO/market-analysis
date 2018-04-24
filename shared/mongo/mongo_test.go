package mongo

import (
	"market-analysis/config"
	"testing"
)

func TestNew(t *testing.T) {
	conf, _ := config.Load("../../config", "config")
	t.Log("conf", conf)
	InitMongo()
}
