package app

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostHandle(t *testing.T) {

	type want struct {
		host        string
		method      string
		body        io.Reader
		code        int
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Try to POST http://yandex.ru",
			want: want{
				host:        "http://localhost:8080",
				method:      "POST",
				body:        strings.NewReader("yandex.ru"),
				code:        201,
				contentType: "text/plain",
			},
		},
		{
			name: "Check true method",
			want: want{
				host:        "http://localhost:8080",
				body:        strings.NewReader("string"),
				method:      "GET",
				code:        400,
				contentType: "",
			},
		},
		{
			name: "Check body",
			want: want{
				host:        "http://localhost:8080",
				body:        strings.NewReader(""),
				method:      "POST",
				code:        201,
				contentType: "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStorage()
			w := httptest.NewRecorder()
			MakePostHandle(store)

			res := w.Result()

			assert.Equal(t, res.StatusCode, tt.want.code)
			defer res.Body.Close()

			assert.Equal(t, res.Header.Get("Content-Type"), tt.want.contentType)
		})
	}
}

func TestGetHandle(t *testing.T) {
	type want struct {
		host   string
		method string
		body   string
		code   int
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Try to GET http://yandex.ru",
			want: want{
				host:   "http://localhost:8080",
				method: "GET",
				body:   "yandex.ru",
				code:   307,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStorage()
			newUrl := "/" + uuid.NewString()
			store.store[newUrl] = tt.want.body

			//request, _ := http.NewRequest(tt.want.method, tt.want.host+newUrl, nil)
			w := httptest.NewRecorder()
			MakeGetHandle(store)

			res := w.Result()

			assert.Equal(t, res.StatusCode, tt.want.code)
			defer res.Body.Close()

		})
	}
}
