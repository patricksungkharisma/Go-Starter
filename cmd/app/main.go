package main

import (
	"log"

	"github.com/patricksungkharisma/go-starter/internal/api/http"
)

func main() {
	err := http.ServerHTTP()
	if err != nil {
		log.Fatalln(err)
	}
}
