package main

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type Logger struct {
	Log logr.Logger
}

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	log := &Logger {
		Log: zapr.NewLogger(logger),
	}


	log.Log.Info("Log example")
}