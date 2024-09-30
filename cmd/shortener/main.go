package main

import (
	"net/http"

	"github.com/fyR27/URL-shortening-service/config"
	"github.com/fyR27/URL-shortening-service/internal/app"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewMux()
	c := config.NewConfig()
	config.ParseFlags(c)
	store := app.NewStore(c)

	r.Post("/", app.MakePostHandle(store))
	r.Get("/{id}", app.MakeGetHandle(store))

	if err := http.ListenAndServe(c.Host, r); err != nil {
		panic(err)
	}
}
