package main

import (
	"log"
	"net/http"
)

func main() {
	router := DecoratedRouter()
 
	log.Fatal(http.ListenAndServe(":8080", router))
}
