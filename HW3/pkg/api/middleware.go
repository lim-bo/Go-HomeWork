package api

import (
	"encoding/base64"
	"net/http"
)

const (
	user1 = "Dwayne:Johnson" //Базу данных может трогать только Скала
)

func (api *api) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.logger.Info("request-info: ", "url", r.URL.Path, "method", r.Method)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH")
		w.Header().Set("Content-Type", "application/json")

		w.Header().Set("WWW-Authenticate", `Basic realm="books-db api", charset="UTF-8"`)
		if r.Method != http.MethodGet {
			ok, err := basicAuth(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			api.logger.Info("login", "user", r.Header.Get("Authorization"))
			if !ok {
				http.Error(w, "Wrong username or password", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})

}

func basicAuth(r *http.Request) (bool, error) {
	auth := r.Header.Get("Authorization")

	source, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return false, err
	}

	return (string(source) == user1), nil
}
