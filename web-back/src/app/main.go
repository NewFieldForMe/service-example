package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Start to serve")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
