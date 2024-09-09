package app

import (
	"fmt"
	"io"
	"net/http"
)

func ParsedBaseURL(host, url, body string) string {
	if url != "" && url != "http://localhost" || ((url[len(url)-1] >= 48 && url[len(url)-1] <= 57) && url != "http://localhost") {
		return url + "/" + body
	} else {
		return url + host + "/" + body
	}
}

func MakePostHandle(store *Storage, host, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		parsedBody := store.AddNewURL(body)
		data := ParsedBaseURL(host, url, parsedBody)
		fmt.Println(data)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)

		if _, err := w.Write([]byte(data)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func MakeGetHandle(store *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := store.FindAddr(r.URL.Path[1:])
		if data == "Bad id" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, data, http.StatusTemporaryRedirect)
	}
}
