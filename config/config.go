package config

import (
	gcfg "gopkg.in/gcfg.v1"

	"github.com/albert-widi/transaction_example/env"
)

// ParseConfig will parse module config
func ParseConfig(cfg interface{}, appName, module string, path ...string) error {
	environ := env.Get()
	return readModuleConfig(cfg, appName, environ, module, path...)
}

// ReadModuleConfig will read config from .ini file
func readModuleConfig(cfg interface{}, appName, env, module string, path ...string) error {
	var err error
	for _, val := range path {
		fname := val + "/" + appName + "/" + module + "." + env + ".ini"
		err = gcfg.ReadFileInto(cfg, fname)
		if err == nil {
			break
		}
	}
	return err
}
