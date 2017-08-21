package env

import (
	"errors"
	"os"
)

const (
	EnvDev   = "dev"
	EnvDebug = "debug"
)

//Get return string of current environment flag
func Get() string {
	env := os.Getenv("TXENV")
	if env == "" {
		env = EnvDebug
	}
	return env
}

func GetAppName() (string, error) {
	appname := os.Getenv("TXAPPNAME")
	if appname == "" {
		return "", errors.New("Application name is empty")
	}
	return appname, nil
}
