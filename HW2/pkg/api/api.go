package api

import (
	"log"
	"net/http"
	"sync"
)

type api struct {
	adress     string
	r          *http.ServeMux
	data       []inputdt
	bght_items []inputdt
	m          *sync.RWMutex
}

type inputdt struct {
	Item  string `json:"item"`
	Count int    `json:"count"`
	Price int    `json:"price"`
}

func New(adress string, r *http.ServeMux) *api {
	new_v := api{adress: adress, r: r}

	var data, bght_items []inputdt
	new_v.data = data
	new_v.bght_items = bght_items

	var mx sync.RWMutex
	new_v.m = &mx

	return &new_v
}

func (api_var *api) FillEndpoints() {
	api_var.r.HandleFunc("/shoppinglist", api_var.shoppingListHandler)
	api_var.r.HandleFunc("/buyitem", api_var.buyItemHandler)
}

func (api_var *api) ListenAndServe() {
	log.Fatal(http.ListenAndServe(api_var.adress, api_var.r))
}

func remove(s []inputdt, i int) []inputdt {
	return append(s[:i], s[i+1:]...)
}

func (api_var *api) buyItem(input inputdt, elInd int) {
	if api_var.data[elInd].Count-input.Count <= 0 {
		input.Count = api_var.data[elInd].Count
		input.Price = api_var.data[elInd].Price
		api_var.bght_items = append(api_var.bght_items, input)
		api_var.data = remove(api_var.data, elInd)
	} else {
		api_var.data[elInd].Count -= input.Count
		input.Price = api_var.data[elInd].Price
		api_var.bght_items = append(api_var.bght_items, input)
	}
}
