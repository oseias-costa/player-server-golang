package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayeStorage interface {
	GetPoints(name string) int
}

type PlayerServer strucut {
	storage PlayeStorage
}

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5001", handler); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, GetPoints(player))

}

func GetPoints(name string) string {
	if name == "Maria" {
		return "20"
	}

	if name == "Pedro" {
		return "10"
	}

	return ""
}
