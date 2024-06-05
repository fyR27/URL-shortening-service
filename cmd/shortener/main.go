package main

import (
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusCreated)

	data := []byte("http://localhost:8080/EwHXdJfB")

	_, err := w.Write(data)
	if err != nil {
		panic(err)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusTemporaryRedirect)

	data := []byte("https://practicum.yandex.ru/")

	_, err := w.Write(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", PostHandler)
	mux.HandleFunc("/EwHXdJfB", GetHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
