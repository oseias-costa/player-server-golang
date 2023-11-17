package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func GetCatFact() {
	url := "https://catfact.ninja/fact"
	var catfact CatFact

	err := GetJson(url, catfact)
	if err != nil {
		fmt.Printf("error getting cat fact: %s\n", err.Error())
		return
	} else {
		fmt.Printf("A super interesting Cat Fact", catfact.Fact)
	}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetCatFact()
}
