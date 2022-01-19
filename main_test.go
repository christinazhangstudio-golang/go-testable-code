// +build unit

package main

import (
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"
	"io/ioutil"
)

type AddResult struct {
	x 			int
	y 			int
	expected 	int
}

var addResults = []AddResult {
	{1, 1, 2},
	//can add more test cases
}

func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if (result != test.expected) {
			t.Fatal("Expected result not given")
		}
	}
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"good\"}")
	}

	//perform the request
	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if 200 != resp.StatusCode {
		t.Fatal("Status Code Not Ok")
	}
}
