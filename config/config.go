package config

import (
	"flag"

	gcfg "gopkg.in/gcfg.v1"

	"github.com/albert-widi/transaction_example/database"
	"github.com/albert-widi/transaction_example/env"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/redis"
)

// This package will load all application config

type AppConfig struct {
	Database database.Config
	Redis    redis.Config
	l        log.Config
}

var directories = []string{
	"files/config/transactionapp",
	"../files/config/transactionapp",
	"../../files/config/transactionapp",
	"../../../files/config/transactionapp",
}

// variables for CLI
var (
	logLevel     = flag.String("log_level", "info", "set log level")
	errorLogPath = flag.String("error_log", "", "log path")
)

func GetConfig() (AppConfig, error) {
	flag.Parse()
	conf := AppConfig{
		l: log.Config{
			LogLevel:     *logLevel,
			ErrorLogPath: *errorLogPath,
		},
	}
	log.SetConfig(conf.l)

	// load database config
	databaseConfig := database.Config{}
	err := ParseConfig(&databaseConfig, "database", directories...)
	if err != nil {
		return conf, err
	}
	conf.Database = databaseConfig

	// load redis config
	redisConfig := redis.Config{}
	err = ParseConfig(&redisConfig, "redis", directories...)
	if err != nil {
		return conf, err
	}
	conf.Redis = redisConfig

	log.Debugf("Application config: %+V", conf)
	return conf, nil
}

// ParseConfig will parse module config
func ParseConfig(cfg interface{}, module string, path ...string) error {
	environ := env.Get()
	return readModuleConfig(cfg, environ, module, path...)
}

// ReadModuleConfig will read config from .ini file
func readModuleConfig(cfg interface{}, env, module string, path ...string) error {
	var err error
	for _, val := range path {
		fname := val + "/" + module + "." + env + ".ini"
		err = gcfg.ReadFileInto(cfg, fname)
		if err == nil {
			break
		}
	}
	return err
}
