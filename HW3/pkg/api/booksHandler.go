package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/pkg/models"
)

func (api *api) books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		id, got := vars["id"]
		if got {
			id, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			b, err := api.database.ReadBook(context.Background(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = w.Write([]byte(fmt.Sprint(b)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		books, err := api.database.ReadBooks(context.Background(), 0, 100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write([]byte(fmt.Sprint(books)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		var b models.Book
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.database.AddBook(context.Background(), b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodDelete:
		vars := mux.Vars(r)
		id, got := vars["id"]
		if !got {
			http.Error(w, "bookId param requied\n", http.StatusInternalServerError)
			return
		}
		var intId int
		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "convertion error\n", http.StatusInternalServerError)
			return
		}
		err = api.database.RemoveBook(context.Background(), intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPatch: //Обновление записи по id, id передается в параметрах строки подключения
		vars := mux.Vars(r)
		id, got := vars["id"]
		if !got {
			http.Error(w, "bookId requied for update\n", http.StatusInternalServerError)
			return
		}

		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "convertion error\n", http.StatusInternalServerError)
			return
		}

		var b models.Book
		err = json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = api.database.UpdateBooksAt(context.Background(), intId, b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
