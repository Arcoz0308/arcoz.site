package main

import (
	"arcoz.dev/api/routes"
	"arcoz.dev/api/utils"
	"log"
	"net/http"
)

var (
	server http.Server
)

func main() {
	utils.LoadConfig()
	server = http.Server{
		Addr:    ":80",
		Handler: routes.Init(),
	}
	log.Fatal(server.ListenAndServe())
}
