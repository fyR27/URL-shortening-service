package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func SendRequestToServer(t *testing.T, code int, method, body string) *http.Response {
	var (
		handler http.HandlerFunc
		req     *http.Request
		store   = NewStore()
	)

	if method == "GET" {
		validID := uuid.NewString()
		if body != "" {
			store.store[validID] = body
		}
		req = httptest.NewRequest(method, "http://localhost:8080/"+validID, nil)
	} else {
		req = httptest.NewRequest(method, "http://localhost:8080/", strings.NewReader(body))
	}
	w := httptest.NewRecorder()

	if method == "POST" {
		handler = MakePostHandle(store)
	} else if method == "GET" {
		handler = MakeGetHandle(store)
	}
	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, code, res.StatusCode)
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
				body:        "",
				method:      "POST",
				code:        http.StatusBadRequest,
				contentType: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SendRequestToServer(t, tt.want.code, tt.want.method, tt.want.body)
			defer res.Body.Close()
			
			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
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
				body:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SendRequestToServer(t, tt.want.code, tt.want.method, tt.want.body)
			defer res.Body.Close()
		})
	}
}
