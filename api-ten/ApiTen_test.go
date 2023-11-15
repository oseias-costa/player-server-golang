package apiten

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiTen(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ApiTen))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expect := "TEN"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expect {
		t.Errorf("expect %s, but got %s", expect, string(b))
	}
}
