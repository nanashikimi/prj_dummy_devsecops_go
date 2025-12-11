package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func FuzzHelloHandler(f *testing.F) {
	f.Add("world")
	f.Add("DevSecOps")

	f.Fuzz(func(t *testing.T, name string) {
		escaped := url.QueryEscape(name)

		req := httptest.NewRequest("GET", "/?name="+escaped, nil)
		w := httptest.NewRecorder()

		helloHandler(w, req)

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("unexpected status code: %d", w.Result().StatusCode)
		}
	})
}
