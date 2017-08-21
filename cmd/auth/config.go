package main

import (
	"errors"
	"flag"

	"github.com/albert-widi/transaction_example/config"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/redis"
)

type AppConfig struct {
	Redis redis.Config
	l     log.Config
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
	appName      = flag.String("appname", "", "set application name")
	logLevel     = flag.String("log_level", "info", "set log level")
	errorLogPath = flag.String("error_log", "", "log path")
)

func ApplicationConfig() (AppConfig, error) {
	flag.Parse()
	conf := AppConfig{
		l: log.Config{
			LogLevel:     *logLevel,
			ErrorLogPath: *errorLogPath,
			AppName:      *appName,
		},
	}
	if *appName == "" {
		return conf, errors.New("Application name cannot be empty")
	}
	log.SetConfig(conf.l)

	// load redis config
	redisConfig := redis.Config{}
	err := config.ParseConfig(&redisConfig, *appName, "redis", directories...)
	if err != nil {
		return conf, err
	}
	conf.Redis = redisConfig

	log.Debugf("Application config: %+V", conf)
	return conf, nil
}
