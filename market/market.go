package market

import (
	"fmt"
	"market-analysis/shared/request"
)

type Market struct {
}

func (m *Market) GetTicker() {

	req := request.New()
	resp, body, errs := request.Get(GetConfig().ApplicationConfiguration.CoinMarketCapURL).End()
	fmt.Println("body", body)
}
