package apitree

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiTree(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ApiTree))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expect := "Treee"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expect {
		t.Errorf("expect %s, response %s", expect, string(b))
	}

}
