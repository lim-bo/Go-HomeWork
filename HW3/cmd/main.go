package main

import (
	"log"

	"github.com/gorilla/mux"
	"main.go/pkg/api"
)

const dbConnString = ""

func main() {
	api := api.New(mux.NewRouter(), dbConnString)
	api.HandleEndpoints()
	log.Fatal(api.ListenAndServe("localhost:8090"))
}
