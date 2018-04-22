package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	conf, _ := Load(".", "config")
	t.Log("TestLoad:config", conf)
	assert.Equal(t, "https://api.coinmarketcap.com", conf.ApplicationConfiguration.CoinMarketCapURL)
}
