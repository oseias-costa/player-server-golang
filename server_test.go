package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("return Maria's result ", func(t *testing.T) {
		req := newReqGetPoints("Maria")
		res := httptest.NewRecorder()

		PlayerServer(res, req)

		verifyReq(t, res.Body.String(), "20")
	})

	t.Run("return Pedro's result", func(t *testing.T) {
		req := newReqGetPoints("Pedro")
		res := httptest.NewRecorder()

		PlayerServer(res, req)

		verifyReq(t, res.Body.String(), "10")
	})

}

func verifyReq(t *testing.T, receiver, expect string) {
	t.Helper()
	if receiver != expect {
		t.Errorf("receiver %s, expect %s", receiver, expect)
	}
}

func newReqGetPoints(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
