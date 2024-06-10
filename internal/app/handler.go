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

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusCreated)

	data := []byte("localhost:8080" + s.AddNewUrl(body))

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
	w.Header().Set("Content-Type", "text/plain")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, s.FindAddr(string(body[:])), http.StatusTemporaryRedirect)
}
