package foohandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expect 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expect := "FOO"
	b, err := io.ReadAll(resp.Body)
	if string(b) != expect {
		t.Errorf("expect %s but we got %s", expect, string(b))
	}

}
