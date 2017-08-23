package main

import (
	"flag"

	"github.com/albert-widi/transaction_example/apicalls"
	"github.com/albert-widi/transaction_example/config"
	"github.com/albert-widi/transaction_example/database"
	"github.com/albert-widi/transaction_example/env"
	"github.com/albert-widi/transaction_example/log"
)

type AppConfig struct {
	Database database.Config
	APICalls apicalls.Config
	l        log.Config
}

var directories = []string{
	"/etc",
	"files/config",
	"../files/config",
	"../../files/config",
	"../../../files/config",
}

// variables for CLI
var (
	logLevel     = flag.String("log_level", "info", "set log level")
	errorLogPath = flag.String("error_log", "", "log path")
)

func ApplicationConfig() (AppConfig, error) {
	flag.Parse()
	conf := AppConfig{
		l: log.Config{
			LogLevel:     *logLevel,
			ErrorLogPath: *errorLogPath,
		},
	}
	appName, err := env.GetAppName()
	if err != nil {
		return conf, err
	}
	conf.l.AppName = appName
	log.SetConfig(conf.l)

	apicallsConfig := apicalls.Config{}
	err = config.ParseConfig(&apicallsConfig, "", "apicalls", directories...)
	if err != nil {
		return conf, err
	}
	conf.APICalls = apicallsConfig

	// load redis config
	dbConfig := database.Config{}
	err = config.ParseConfig(&dbConfig, appName, "database", directories...)
	if err != nil {
		return conf, err
	}
	conf.Database = dbConfig

	log.Debugf("Application config: %+V", conf)
	return conf, nil
}
