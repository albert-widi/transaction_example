package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/albert-widi/transaction_example/cmd/logistic/logisticapi"
	"github.com/albert-widi/transaction_example/database"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/service/httpapi"
)

func main() {
	// flag is parsed in config, don't parse it anyhwere else
	appConfig, err := ApplicationConfig()
	if err != nil {
		log.Fatal("Failed to get aplication config: ", err.Error())
	}
	// connect to redis
	err = database.Init(appConfig.Database)
	if err != nil {
		log.Fatal("Failed to init database ", err.Error())
	}

	fatalChan := make(chan error)
	w := httpapi.New(httpapi.Config{
		ListenAddress:  ":9003",
		RouteEndpoints: logisticapi.Endpoints,
	})
	go func() {
		if err := w.Run(); err != nil {
			fatalChan <- err
		}
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Signal terminate detected")
	case err := <-fatalChan:
		log.Fatal("Application failed to run because ", err.Error())
	}
	log.Warn("Logisticapp exited...")
}
