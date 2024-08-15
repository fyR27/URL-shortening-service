package main

import (
	"net/http"

	"github.com/fyR27/URL-shortening-service/internal/app"
)

func main() {
	store := app.NewStore()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", app.MakePostHandle(store))
	mux.HandleFunc("GET /{id}", app.MakeGetHandle(store))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
