package main

import (
	"net/http"

	"github.com/fyR27/URL-shortening-service/config"
	"github.com/fyR27/URL-shortening-service/internal/app"
	"github.com/go-chi/chi/v5"
)

func main() {
	c := config.NewConfig()
	config.ParseFlags(c)
	r := chi.NewMux()
	store := app.NewStore()

	r.Post("/", app.MakePostHandle(store, c.Host, c.URL))
	r.Get("/{id}", app.MakeGetHandle(store))

	if err := http.ListenAndServe(c.Host, r); err != nil {
		panic(err)
	}
}
