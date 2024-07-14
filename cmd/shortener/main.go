package main

import (
	"net/http"

	app "github.com/fyR27/URL-shortening-service/internal/app"
)

func main() {
	store := new(app.Storage)
	store = app.NewStore()
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.MakePostHandle(store))
	mux.HandleFunc("GET /{id}", app.MakeGetHandle(store))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
