package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	testCases := []struct {
		URL    string
		Expect string
	}{

		{"http://example.com/", "Not Found\n"},
		{"http://example.com/goodbye", "Not Found\n"},
		{"http://example.com/hello", "Not Found\n"},
		{"http://example.com/hello/", "Not Found\n"},
		{"http://example.com/hello/luke", "Hello, Luke!\n"},
		{"http://example.com/hello/mcWinner", "Hello, McWinner!\n"},
		{"http://example.com/hello/Doe", "Hello, Doe!\n"},
		{"http://example.com/hello/james%20bond", "Hello, James Bond!\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.URL+" == "+tc.Expect, func(subt *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tc.URL, nil)
			root(w, req)
			resp := w.Result()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("ReadAll(): %v", err)
			}
			actual := string(body)
			if actual != tc.Expect {
				t.Errorf("body == '%s'", actual)
			}
		})
	}
}
