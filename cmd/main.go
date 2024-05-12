package main

import (
	"log"

	"github.com/RadekKusiak71/goEcom/cmd/api"
	"github.com/RadekKusiak71/goEcom/db"
)

func main() {
	storage, err := db.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", storage)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
