package market

import (
	"fmt"
	"market-analysis/config"
	"market-analysis/shared/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Print("market_test:init")
	config.Load("../config", "config")
	log.InitLogger()
}

func TestGetTicker(t *testing.T) {
	tickers, _ := (&Market{}).GetTicker()
	assert.NotNil(t, tickers)

}
