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

	return res
}

func TestPostHandle(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}

	tests := []struct {
		name   string
		method string
		body   string
		want   want
	}{
		{
			name:   "Try to POST http://yandex.ru",
			method: "POST",
			body:   "http://yandex.ru",
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain",
			},
		},
		{
			name:   "Check empty body",
			method: "POST",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore()
			handler := MakePostHandle(store)

			res := SendRequestToServer(t, &handler, "http://localhost:8080/", tt.method, tt.body)
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			if res.StatusCode != http.StatusCreated {
				assert.Equalf(t, res.Header.Get("Content-Type"), tt.want.contentType, "Content type %s is not equal to %s")
			} else {
				assert.Equal(t, res.Header.Get("Content-Type"), tt.want.contentType)
			}

		})
	}
}

func TestGetHandle(t *testing.T) {
	type want struct {
		code int
	}

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		want   want
	}{
		{
			name:   "Try to GET with valid ID",
			method: "GET",
			path:   "/get/invalid",
			body:   "http://yandex.ru",
			want: want{
				code: http.StatusTemporaryRedirect,
			},
		},
		{
			name:   "Try to GET with invalid ID",
			method: "GET",
			path:   "get/valid",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore()
			validID := uuid.NewString()

			if tt.body != "" {
				store.store[validID] = tt.body
			}

			handler := MakeGetHandle(store)

			res := SendRequestToServer(t, &handler, "http://localhost:8080/"+validID, tt.method, uuid.Nil.String())
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
