package main

import (
	log "github.com/sirupsen/logrus"
	"golang.frontdoorhome.com/personal-project/tiny-url/internal/server"
)

func main() {
	err := server.RunServer()
	if err != nil {
		log.Fatal(err)
	}
}
