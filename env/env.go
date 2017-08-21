package env

import "os"

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

const appname = "transactionapp"

func GetAppName() string {
	return appname
}
