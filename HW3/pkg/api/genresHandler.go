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

func (api *api) genres(w http.ResponseWriter, r *http.Request) {
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
			g, err := api.database.ReadGenre(context.Background(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = w.Write([]byte(fmt.Sprint(g)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		genres, err := api.database.ReadGenres(context.Background(), 0, 100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write([]byte(fmt.Sprint(genres)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		var g models.Genre
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.database.AddGenre(context.Background(), g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodDelete:
		vars := mux.Vars(r)
		id, got := vars["id"]
		if !got {
			http.Error(w, "genreId param requied\n", http.StatusInternalServerError)
			return
		}
		var intId int
		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "convertion error\n", http.StatusInternalServerError)
			return
		}
		err = api.database.RemoveGenre(context.Background(), intId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPatch: //Обновление записи по id, id передается в параметрах строки подключения
		vars := mux.Vars(r)
		id, got := vars["id"]
		if !got {
			http.Error(w, "genreId requied for update\n", http.StatusInternalServerError)
			return
		}

		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "convertion error\n", http.StatusInternalServerError)
			return
		}

		var g models.Genre
		err = json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = api.database.UpdateGenresAt(context.Background(), intId, g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
