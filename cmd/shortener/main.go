package main

import (
	"net/http"

	"github.com/fyR27/URL-shortening-service/internal/app"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()
	store := app.NewStore()
	r.Post("/", app.MakePostHandle(store))
	r.Get("/{id}", app.MakeGetHandle(store))

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
