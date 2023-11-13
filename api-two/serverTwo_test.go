package apitwo

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiTwo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ApitTwo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expect := "Api Two"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expect {
		t.Errorf("expected %s but we got %s", expect, string(b))
	}

}
