package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func init() {
	gin.SetMode(gin.TestMode)
	r.GET("/ping", Ping)
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
				expectedBody:         "{\"message\":\"Welcome to the Coffeeshop!\"}",
			},
		},
	}
	for _, tt := range tests {
		req := getRequest(tt.args.method, tt.args.path)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		body, _ := ioutil.ReadAll(w.Body)
		if tt.want.expectedBody != string(body) {
			t.Errorf("Expected '%v', but got '%v'", tt.want.expectedBody, string(body))
		}

		if tt.want.expectedResponseCode != w.Code {
			t.Errorf("Expected '%v', but got '%v'", tt.want.expectedResponseCode, w.Code)
		}
		// t.Run(tt.name, func(t *testing.T) {
		// 	Ping(tt.args.c)
		// })
	}
}

func getRequest(method, path string) *http.Request {
	req, _ := http.NewRequest("GET", "/ping", nil)
	return req
}
