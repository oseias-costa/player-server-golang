package main

import (
	"foohandler/unit-test-k8s/api"
	"log"
	"net/http"
)

func main() {
	service := api.Service{
		HTTPClient: http.DefaultClient,
	}

	http.HandleFunc("/", service.GetRandomUser)
	log.Println("Listening...")
	http.ListenAndServe(":7000", nil)
}
