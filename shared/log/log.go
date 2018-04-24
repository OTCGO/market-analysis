package log

import (
	"encoding/json"
	"fmt"
	"market-analysis/config"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() {
	configration, _ := config.GetConfig()
	rawJSON := fmt.Sprintf(`{
		"level": "%s",
		"encoding": "json",
		"outputPaths": ["stdout","%s"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		  }
		}`, configration.LogConfiguration.Level, configration.LogConfiguration.Output)
	var cfg zap.Config
	if err := json.Unmarshal([]byte(rawJSON), &cfg); err != nil {
		errors.Wrap(err, "new New json.Unmarshal error")
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	logger, err = cfg.Build()
	if err != nil {
		errors.Wrap(err, "log build error")
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")
}

func GetLogger() *zap.Logger {
	if logger == nil {
		return nil
	}
	return logger
}
