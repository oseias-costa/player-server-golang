package api

import (
	"io"
	"log"
	"net/http"
)

type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type Service struct {
	HTTPClient HTTPClientInterface
}

func (s *Service) GetRandomUser(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://randomuser.me/api"
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Println(err)
	}

	res, err := s.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
