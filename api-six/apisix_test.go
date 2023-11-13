package apisix

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiSix(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ApiSix))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expect := "sixxx"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if expect != string(b) {
		t.Errorf("expect %s, but got %s", expect, string(b))
	}
}
