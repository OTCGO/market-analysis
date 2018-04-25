package main

import (
	"fmt"
	"market-analysis/config"
	"market-analysis/market"
	"market-analysis/model"
	"market-analysis/shared/log"
	"market-analysis/shared/mongo"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"go.uber.org/zap"
)

var logger *zap.Logger

var maxRoutineNum int = 10

var page = 50
var pageCount int = 10
var configration *config.Config

func init() {
	configration, _ = config.Load("./config", "config")
	log.InitLogger()
	mongo.InitMongo()
	logger = log.GetLogger()
}

func main() {

	logger.Info("start")
	err := (&model.Ticker{}).CreateIndex()
	if err != nil {
		errors.Wrap(err, "ticker create index serror")
	}

	c := cron.New()
	// second minute hour day_of_month month day_of_week
	c.AddFunc(configration.ApplicationConfiguration.CronTime, func() {
		fmt.Println("Every minute", time.Now().Format("2006-01-02 15:04:05"))
		startup()
	})

	c.Start()

	select {}
}

//result channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		tickers, err := (&market.Market{}).GetTicker(page*(j-1), page)

		if err != nil {
			errors.Wrap(err, "request ticker serror")
		}

		for _, ticker := range tickers {
			err := ticker.Insert()
			if err != nil {
				errors.Wrap(err, "ticker Insert serror")
			}
		}

		results <- j
	}
}

func startup() {
	// make two channel。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// startup
	for w := 1; w <= maxRoutineNum; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= pageCount; j++ {
		jobs <- j
	}
	close(jobs)

	//handle result
	for a := 1; a <= pageCount; a++ {
		<-results
	}
}
