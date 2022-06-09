package main

import (
	"log"
	"net/http"

	"github.com/go-web-code/rest-go/routes"
)

func main() {
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
