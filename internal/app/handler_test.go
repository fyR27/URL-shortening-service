package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func SendRequestToServer(t *testing.T, handler *http.HandlerFunc, url, method, body string) *http.Response {

	req := httptest.NewRequest(method, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	res := w.Result()
	res.Body.Close()

	return res
}

func TestPostHandle(t *testing.T) {
	type want struct {
		method      string
		body        string
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
				method:      "POST",
				body:        "http://yandex.ru",
				code:        http.StatusCreated,
				contentType: "text/plain",
			},
		},
		{
			name: "Check empty body",
			want: want{
				method: "POST",
				code:   http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore()
			handler := MakePostHandle(store)
			res := SendRequestToServer(t, &handler, "http://localhost:8080/", tt.want.method, tt.want.body)

			assert.Equal(t, tt.want.code, res.StatusCode)
			if tt.want.contentType != "text/plain" {
				assert.Equalf(t, res.Header.Get("Content-Type"), tt.want.contentType, "Content type %s is not equal to %s")
			}
		})
	}
}

func TestGetHandle(t *testing.T) {
	type want struct {
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
				method: "GET",
				path:   "/get/invalid",
				code:   http.StatusTemporaryRedirect,
				body:   "http://yandex.ru",
			},
		},
		{
			name: "Try to GET with invalid ID",
			want: want{
				method: "GET",
				path:   "get/valid",
				code:   http.StatusBadRequest,
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
			handler := MakeGetHandle(store)
			res := SendRequestToServer(t, &handler, "http://localhost:8080/"+validID, tt.want.method, uuid.Nil.String())

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
