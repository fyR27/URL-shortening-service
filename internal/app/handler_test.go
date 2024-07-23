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
				body:        strings.NewReader("http://yandex.ru"),
				code:        201,
				contentType: "text/plain",
			},
		},
		{
			name: "Check empty body",
			want: want{
				host:        "http://localhost:8080",
				body:        strings.NewReader(""),
				method:      "POST",
				code:        400,
				contentType: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore()
			req := httptest.NewRequest(tt.want.method, tt.want.host, tt.want.body)
			w := httptest.NewRecorder()

			handler := MakePostHandle(store)
			handler.ServeHTTP(w, req)

			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			defer res.Body.Close()

			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
			}
		})
	}
}

func TestGetHandle(t *testing.T) {
	type want struct {
		host   string
		method string
		path   string
		code   int
		body   string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Try to GET with valid ID",
			want: want{
				host:   "http://localhost:8080/",
				method: "GET",
				path:   "/get/invalid",
				code:   307,
				body:   "http://yandex.ru",
			},
		},
		{
			name: "Try to GET with invalid ID",
			want: want{
				host:   "http://localhost:8080/",
				method: "GET",
				path:   "get/valid",
				code:   400,
				body:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore()
			validID := uuid.NewString()
			if tt.want.body != "" {
				store.store[validID] = tt.want.body
			}

			req := httptest.NewRequest(tt.want.method, tt.want.host+validID, nil)
			w := httptest.NewRecorder()

			handler := MakeGetHandle(store)
			handler.ServeHTTP(w, req)

			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			defer res.Body.Close()
		})
	}
}
