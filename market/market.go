package market

import (
	"encoding/json"
	"fmt"
	"market-analysis/config"
	"market-analysis/model"
	"market-analysis/shared/log"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Market struct {
}

var conf *config.Config
var logger *zap.Logger

// func init() {

// }

func (m *Market) GetTicker(start int, limit int) (tickers []*model.Ticker, err error) {
	conf, _ = config.GetConfig()
	logger = log.GetLogger()
	logger.Info("init:conf", zap.String("conf", fmt.Sprintf("%v", conf)))

	var url string = fmt.Sprintf("%s/%s?%s", conf.ApplicationConfiguration.CoinMarketCapURL, "v1/ticker", fmt.Sprintf("convert=CNY&start=%d&limit=%d", start, limit))

	logger.Info("GetTicker:url", zap.String("url", url))

	_, body, _ := gorequest.New().Get(url).End()

	// logger.Info("GetTicker:body", zap.String("body", body))
	fmt.Println("body", body)
	// fmt.Println("rs", rs)

	// tickers = make([]model.Ticker, 0)
	// result := map[string]string{"body": body}
	tickers = []*model.Ticker{}
	// fmt.Println("value", result)
	err = json.Unmarshal([]byte(body), &tickers)
	if err != nil {
		errors.Wrap(err, "json.Unmarshal error")
	}
	logger.Info("GetTicker:tickers", zap.String("tickers", fmt.Sprintf("%v", len(tickers))))
	return tickers, err
}
