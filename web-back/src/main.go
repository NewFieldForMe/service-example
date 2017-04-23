package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", 1}
	js, err := json.Marshal(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	log.Printf("Start to serve")
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe("localhost:8000", nil)
}
