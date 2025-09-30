package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initalDB()
	serveHTTP()
}

func initalDB() {
	fmt.Println("Connecting to database")

	db, _ := sql.Open("sqlite3", "recipe.db")
	defer db.Close()
	db.Exec(`PRAGMA journal_mode=WAL`)

	db.Exec(`CREATE TABLE recipes (
		id INTEGER PRIMARY KEY,
		name TEXT NOT_NULL,
		recipe TEXT NOT_NULL,
		created DATETIME DEFAULT CURRENT_TIMESTAMP,
		favorite BOOL)
		`)

	db.Exec(`INSERT INTO recipes (name, recipe) VALUES ('Pancakes', '3 eggs, 3 dl flour, 5 dl milk, butter')`)

	fmt.Println("Database operations complete")
}

func serveHTTP() {
	mux := http.NewServeMux()

	log.Print("Listening on http://localhost:3000")
	mux.HandleFunc("GET /recipe/all}", allRecipes)

	http.ListenAndServe(":3000", mux)
}

func allRecipes(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "recipe.db")
	recipes, err := db.Exec(`SELECT * FROM recipes`)
	if err != nil {
		error.Error(err)
	}
	fmt.Sprintf("%v", recipes)
}
