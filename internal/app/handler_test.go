package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fyR27/URL-shortening-service/config"
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
		name string
		host string
		url  string
		body string

		want want
	}{
		{
			name: "Try to POST http://yandex.ru",
			host: ":8080",
			url:  "http://local",
			body: "http://yandex.ru",

			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain",
			},
		},
		{
			name: "Check empty body",
			host: ":8081",
			url:  "http://localhost",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore()
			c := config.NewConfig()
			c.Host, c.URL = tt.host, tt.url
			handler := MakePostHandle(store, c)

			fmt.Println(handler)

			res := SendRequestToServer(t, &handler, tt.url+tt.host+"/", http.MethodPost, tt.body)
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			if res.StatusCode != http.StatusCreated {
				assert.Equal(t, res.Header.Get("Content-Type"), tt.want.contentType)
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
		name string
		path string
		body string
		want want
	}{
		{
			name: "Try to GET with valid ID",
			path: "/get/invalid",
			body: "http://yandex.ru",
			want: want{
				code: http.StatusTemporaryRedirect,
			},
		},
		{
			name: "Try to GET with invalid ID",
			path: "get/valid",
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

			res := SendRequestToServer(t, &handler, "http://localhost:8080/"+validID, http.MethodGet, uuid.Nil.String())
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
		})
	}
}
