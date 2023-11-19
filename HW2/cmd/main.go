package main

import (
	"cmd/main/pkg/api"
	"net/http"
)

func main() {
	api_var := api.New("localhost:8080", &http.ServeMux{})
	api_var.FillEndpoints()
	api_var.ListenAndServe()
}
