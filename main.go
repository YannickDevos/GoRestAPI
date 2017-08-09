package main

import (
	"log"
	"net/http"

	"github.com/ricardo-ch/GoRestAPI/restapi"
)

func main() {

	router := restapi.NewRouter1()

	log.Fatal(http.ListenAndServe(":8080", router))
}
