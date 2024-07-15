package app

import (
	"io"
	"net/http"
)

func MakePostHandle(store *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		data := []byte("http://localhost:8080/" + store.AddNewURL(body))

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)

		if _, err := w.Write(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func MakeGetHandle(store *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data := store.FindAddr(r.URL.Path[1:])
		http.Redirect(w, r, data, http.StatusTemporaryRedirect)
	}
}
