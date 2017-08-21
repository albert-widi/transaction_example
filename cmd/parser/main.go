package main

import (
	"log"
	"reflect"

	"github.com/albert-widi/transaction_example/database"
)

type AppConfig struct {
	Database database.Config `config:"database"`
}

var directories = []string{
	"/etc/transactionapp",
	"files/config/transactionapp",
	"../files/config/transactionapp",
	"../../files/config/transactionapp",
	"../../../files/config/transactionapp",
}

func main() {
	app := &AppConfig{}
	ref := reflect.ValueOf(app)
	// ref := reflect.TypeOf(app)
	fieldNum := ref.Len()
	log.Println("NUM: ", fieldNum)

	for x := 0; x < fieldNum; x++ {
		// tag := ref.Field(x).Tag.Get("config")
		// log.Println(tag)
		// log.Printf("%+v", reflect.ValueOf(app).Field(x).Interface())
		// err := config.ParseConfig(reflect.ValueOf(app).Field(x), tag, directories...)
		// if err != nil {
		// 	log.Println("Failed to parse ", err.Error())
		// }
	}
}
