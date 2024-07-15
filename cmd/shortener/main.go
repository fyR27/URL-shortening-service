package main

import (
	"URL-shortening-service/internal/app"
	"net/http"
)

func main() {
	store := app.NewStorage()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", app.MakePostHandle(store))
	mux.HandleFunc("GET /{id}", app.MakeGetHandle(store))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
