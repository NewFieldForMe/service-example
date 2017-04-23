package main

import (
	"log"
	"net/http"
)

type Page struct {
	Title string
	Count int
}

func main() {
	log.Printf("Start to serve")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
