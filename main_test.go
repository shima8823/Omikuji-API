package main_test

import (
	"github.com/shima8823/Omikuji-API"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	main.Handler(w, r)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("unexpected error ", err)
	}
	if len(b) == 0 {
		t.Fatalf("unexpected response: %s", string(b))
	}
}
