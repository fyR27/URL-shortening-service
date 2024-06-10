package app

import (
	"io"
	"net/http"
)

func PostHandle(w http.ResponseWriter, r *http.Request) {
	var s Storage
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	data := []byte("localhost:8080" + string(s.AddNewUrl(body)))

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write(data); err != nil {
		panic(err)
	}
}

func GetHandle(w http.ResponseWriter, r *http.Request) {
	var s Storage
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := s.FindAddr("/EwHXdJfB")

	w.Header().Set("Content-Type", "text/plain")

	http.Redirect(w, r, data, http.StatusTemporaryRedirect)
}
