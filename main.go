package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/ItsAchance/what2eat/handlers"
	_ "github.com/mattn/go-sqlite3"
)

type Recipe struct {
	ID       int
	Name     string
	Recipe   string
	Created  time.Time
	Favorite bool
}

func queryAll(db *sql.DB) []Recipe {
	query := "SELECT id, name, recipe, favorite FROM recipes WHERE DELETED = 0;"
	var id int
	var name string
	var recipe string
	var favorite bool

	data := []Recipe{}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &recipe, &favorite)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Recipe{
			ID:       id,
			Name:     name,
			Recipe:   recipe,
			Favorite: favorite,
		})
	}

	return data
}

func main() {
	// Connect to DB
	db, err := sql.Open("sqlite3", "file:recipe.db?_journal_mode=WAL&_synchronous=NORMAL")
	if err != nil {
		log.Printf("An error has occured: %s\n", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Startpage)

	log.Print("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}
