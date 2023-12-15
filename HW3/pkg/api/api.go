package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/pkg/repository"
)

type api struct {
	database *repository.PgRepo
	r        *mux.Router
}

func New(r *mux.Router, dbConnString string) *api {
	db, err := repository.New(context.Background(), dbConnString)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &api{r: r, database: db}
}

func (api *api) HandleEndpoints() {
	booksRoute := api.r.HandleFunc("/api/books", api.books)
	booksRoute.Queries("id", "{id}")
	booksRoute = api.r.HandleFunc("/api/books", api.books)

	genresRoute := api.r.HandleFunc("/api/genres", api.genres)
	genresRoute.Queries("id", "{id}")
	genresRoute = api.r.HandleFunc("/api/genres", api.genres)

	authorsRoute := api.r.HandleFunc("/api/authors", api.authors)
	authorsRoute.Queries("id", "{id}")
	authorsRoute = api.r.HandleFunc("/api/authors", api.authors)
}

func (api *api) ListenAndServe(adressStr string) error {
	return http.ListenAndServe(adressStr, api.r)
}
