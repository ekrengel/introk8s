package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)


func Test_handlers(t *testing.T) {
	var tests = []struct{
		setup func()
		handler func(http.ResponseWriter, *http.Request)
		expectedBody string
	}{
		{
			setup: func() {},
			handler: hello,
			expectedBody: "Hello ACT-W!",
		},
		{
			setup: func() {},
			handler: kubernetes,
			expectedBody: "Pod name: UNKNOWN\nNode name: UNKNOWN\n",
		},
		{
			setup: func() {
				os.Setenv("POD_NAME", "testPodName")
				os.Setenv("NODE_NAME", "testNodeName")
			},
			handler: kubernetes,
			expectedBody: "Pod name: testPodName\nNode name: testNodeName\n",
		},
	}

	for _, tt := range tests {
		tt.setup()
		req, err := http.NewRequest("GET", "/test", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(tt.handler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if rr.Body.String() != tt.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), tt.expectedBody)
		}
	}
}