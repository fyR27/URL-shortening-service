package app

import (
	"fmt"
	"io"
	"net/http"
)

func PostHandle(w http.ResponseWriter, r *http.Request) {
	var p ParserURL
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := []byte("http://localhost:8080" + p.AddNewURL(body))

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetHandle(w http.ResponseWriter, r *http.Request) {
	var p ParserURL
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := p.FindAddr(r.URL.Path)
	fmt.Println(data)

	http.Redirect(w, r, data, http.StatusTemporaryRedirect)
}
