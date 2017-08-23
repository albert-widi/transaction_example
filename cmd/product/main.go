package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/albert-widi/transaction_example/cmd/product/productapi"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/service/httpapi"
)

func main() {
	fatalChan := make(chan error)
	w := httpapi.New(httpapi.Config{
		ListenAddress:  ":9001",
		RouteEndpoints: productapi.Endpoints,
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
	log.Warn("Productapp exited...")
}
