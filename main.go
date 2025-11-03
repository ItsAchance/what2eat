package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()

	log.Print("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}

