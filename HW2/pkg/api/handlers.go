package api

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
)

func (api_var *api) shoppingListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: //GET-запрос передает в подключение список добавленных продуктов
		if len(api_var.data) != 0 {
			api_var.m.RLock()
			defer api_var.m.RUnlock()

			sum := 0
			for i, item := range api_var.data {
				item_description := strconv.Itoa(i+1) + ". " + item.Item + " count: " + strconv.Itoa(item.Count) + " price: " + strconv.Itoa(item.Price) + "\n"
				_, err := w.Write([]byte(item_description))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				sum += item.Price * item.Count
			}
			list_result := "result: " + strconv.Itoa(sum) + "\n"
			w.Write([]byte(list_result))

		} else {
			_, err := w.Write([]byte("empty list\n"))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	case http.MethodPost: //POST-запрос передает в список продуктов json формата {item_name, count, price}
		api_var.m.Lock()
		defer api_var.m.Unlock()

		var input inputdt
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message := input.Item + " successfuly added to shoppinglist\n"
		_, err = w.Write([]byte(message))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		api_var.data = append(api_var.data, input)

	}
}

func (api_var *api) buyItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: //GET-запрос передает список купленных продуктов
		api_var.m.RLock()
		defer api_var.m.RUnlock()

		sum := 0
		if len(api_var.bght_items) != 0 {
			for i, item := range api_var.bght_items {
				item_description := strconv.Itoa(i+1) + ". " + item.Item + " count: " + strconv.Itoa(item.Count) + " price: " + strconv.Itoa(item.Price) + "\n"
				_, err := w.Write([]byte(item_description))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				sum += item.Count * item.Price
			}
			list_result := "result: " + strconv.Itoa(sum) + "\n"
			w.Write([]byte(list_result))
		} else {
			_, err := w.Write([]byte("You haven't purchased anything yet\n"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	case http.MethodPost: //POST-запрос "совершает" покупку позиции из списка, принимает json формата {item_name, count}
		api_var.m.Lock()
		defer api_var.m.Unlock()

		var input inputdt
		dec := json.NewDecoder(r.Body)
		dec.Decode(&input)
		i := slices.IndexFunc(api_var.data, func(c inputdt) bool { return c.Item == input.Item })
		api_var.buyItem(input, i)
		_, err := w.Write([]byte("Item has bought\n"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
