package apifour

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiFour(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(ApiFour))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expect := "4"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(b) != expect {
		t.Errorf("expect %s, but got %s", expect, string(b))
	}
}
