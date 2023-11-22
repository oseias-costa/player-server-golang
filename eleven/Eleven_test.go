package eleven

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEleven(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Eleven))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	expect := "11"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if expect != string(b) {
		t.Errorf("expect %s, but got %s", expect, string(b))
	}

}
