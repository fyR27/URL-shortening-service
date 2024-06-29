package main

import (
	"URL-shortening-service/internal/app"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.PostHandle)
	mux.HandleFunc("GET /{id}", app.GetHandle)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
