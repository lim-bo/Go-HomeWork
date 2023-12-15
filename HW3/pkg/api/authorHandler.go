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

func (api *api) authors(w http.ResponseWriter, r *http.Request) {
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
			a, err := api.database.ReadAuthor(context.Background(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = w.Write([]byte(fmt.Sprint(a)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		authors, err := api.database.ReadAuthors(context.Background(), 0, 100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write([]byte(fmt.Sprint(authors)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		var a models.Author
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.database.AddAuthor(context.Background(), a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodDelete:
		vars := mux.Vars(r)
		id, got := vars["id"]
		if !got {
			http.Error(w, "authorId param requied\n", http.StatusInternalServerError)
			return
		}
		var intId int
		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "convertion error\n", http.StatusInternalServerError)
			return
		}
		err = api.database.RemoveAuthor(context.Background(), intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPatch: //Обновление записи по id, id передается в параметрах строки подключения
		vars := mux.Vars(r)
		id, got := vars["id"]
		if !got {
			http.Error(w, "authorId requied for update\n", http.StatusInternalServerError)
			return
		}

		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "convertion error\n", http.StatusInternalServerError)
			return
		}

		var a models.Author
		err = json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = api.database.UpdateAuthorsAt(context.Background(), intId, a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
