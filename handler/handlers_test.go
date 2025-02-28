package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func init() {
	gin.SetMode(gin.TestMode)
	r.GET("/ping", Ping)
	r.GET("/coffee", CoffeeDetails)
}

func TestPing(t *testing.T) {
	type args struct {
		method string
		path   string
	}
	type want struct {
		expectedResponseCode int
		expectedBody         string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test Ping",
			args: args{
				method: "GET",
				path:   "/ping",
			},
			want: want{
				expectedResponseCode: http.StatusOK,
				expectedBody: "{\"message\":\"Welcome to the Coffeeshop!\"}",
			},
		},
	}
	for _, tt := range tests {
		req := getRequest(tt.args.method, tt.args.path)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		body, _ := io.ReadAll(w.Body)
		if tt.want.expectedBody != string(body) {
			t.Errorf("Expected '%v', but got '%v'", tt.want.expectedBody, string(body))
		}

		if tt.want.expectedResponseCode != w.Code {
			t.Errorf("Expected '%v', but got '%v'", tt.want.expectedResponseCode, w.Code)
		}
	}
}

func getRequest(method, path string) *http.Request {
	req, _ := http.NewRequest(method, path, nil)
	return req
}

func TestCoffeeDetails(t *testing.T) {
	type args struct {
		method string
		path   string
	}
	type want struct {
		expectedResponseCode int
		expectedBody         string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test GetCoffees",
			args: args{
				method: "GET",
				path:   "/coffee",
			},
			want: want{
				expectedResponseCode: http.StatusOK,
				expectedBody:         "&{[{Latte %!s(float32=3)} {Cappuccino %!s(float32=2.75)} {Flat White %!s(float32=2.25)}]}",
			},
		},
	}
	for _, tt := range tests {
		req := getRequest(tt.args.method, tt.args.path)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		body, _ := io.ReadAll(w.Body)
		if tt.want.expectedBody != string(body) {
			t.Errorf("Expected '%v', but got '%v'", tt.want.expectedBody, string(body))
		}

		if tt.want.expectedResponseCode != w.Code {
			t.Errorf("Expected '%v', but got '%v'", tt.want.expectedResponseCode, w.Code)
		}
	}
}
