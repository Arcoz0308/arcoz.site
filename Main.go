package main

import (
	"arcoz.dev/api/Utils"
	"arcoz.dev/api/routes"
	"log"
	"net/http"
)

var (
	server http.Server
)

func main() {
	Utils.LoadConfig()
	server = http.Server{
		Addr:    ":80",
		Handler: routes.Init(),
	}
	log.Fatal(server.ListenAndServe())
}
