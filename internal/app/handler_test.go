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
			name: "Check true method",
			want: want{
				host:        "http://localhost:8080",
				body:        strings.NewReader("http://string.ru"),
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
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "Try to GET with invalid ID",
			want: want{
				host:   "http://localhost:8080",
				method: "GET",
				path:   "/get/invalid",
				code:   307,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore()
			validID := uuid.NewString()
			store.store[validID] = "http://yandex.ru"

			req := httptest.NewRequest(tt.want.method, tt.want.host+tt.want.path+validID, nil)
			if tt.name == "Try to GET with invalid ID" {
				req = httptest.NewRequest(tt.want.method, tt.want.host+tt.want.path+"invalid", nil)
			}
			w := httptest.NewRecorder()

			handler := MakeGetHandle(store)
			handler.ServeHTTP(w, req)

			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			defer res.Body.Close()
		})
	}
}
