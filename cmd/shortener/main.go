package main

import (
	"net/http"

	app "github.com/fyR27/URL-shortening-service/internal/app"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.PostHandle)
	mux.HandleFunc("/{id}", app.GetHandle)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
