package log

import (
	"market-analysis/config"
	"testing"
)

func TestLog(t *testing.T) {
	conf, _ := config.Load("../../config", "config")
	t.Log("conf", conf)
	InitLogger()
	GetLogger().Info("TestLog")
}
