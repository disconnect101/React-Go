package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080
const domain = "drumadi.in"

type Application struct {
	Domain string
}

func main() {
	var app Application

	app.Domain = domain
	addr := fmt.Sprintf(":%d", port)

	log.Printf("Starting Application Server on port: %d", port)
	err := http.ListenAndServe(addr, app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
