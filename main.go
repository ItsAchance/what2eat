package main

import (
	"log"
	"net/http"

	"github.com/ItsAchance/what2eat/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Startpage)

	log.Print("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}
