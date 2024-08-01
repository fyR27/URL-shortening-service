package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func SendRequestToServer(t *testing.T, url, method, body string) (*httptest.ResponseRecorder, *http.Request) {

	req := httptest.NewRequest(method, url, strings.NewReader(body))
	w := httptest.NewRecorder()

	return w, req
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
			w, req := SendRequestToServer(t, "http://localhost:8080/", tt.want.method, tt.want.body)

			handler := MakePostHandle(store)
			handler.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			if tt.want.contentType != "text/plain" {
				assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
				fmt.Print("Content type is not equal to text/plain \n")
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
			w, req := SendRequestToServer(t, "http://localhost:8080/"+validID, tt.want.method, uuid.Nil.String())
			hanler := MakeGetHandle(store)
			hanler.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
