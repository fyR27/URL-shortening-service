package app

import (
	"io"
	"net/http"
)

func MakePostHandle(store *Storage, host, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil || string(body) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id := store.AddNewURL(body)

		// Формирование итогового сокращённого URL
		parsedBaseURL, err := url.Parse(baseURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Установка пути в сокращённый URL
		parsedBaseURL.Path = path.Join(parsedBaseURL.Path, id)

		// Формирование итоговой строки URL
		shortURL := parsedBaseURL.String()

		// Отправка ответа
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write([]byte(shortURL)); err != nil {
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
