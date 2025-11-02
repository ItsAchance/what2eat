package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if dbExists() == false {
		createDb()
	}
	insertIntoDb()
	allRecipes()
	serveHTTP()
}

func dbExists() bool {
	_, err := os.Stat("recipe.db")
	if err != nil {
		fmt.Printf("Database does not exist\nCreating database\n")
		return false
	}
	return true
}

func createDb() {
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

	fmt.Println("Database created")
}

// This is bad practise, should only open once
func insertIntoDb() {
	db, _ := sql.Open("sqlite3", "recipe.db")
	defer db.Close()
	db.Exec(`INSERT INTO recipes (name, recipe) VALUES ('Pancakes', '3 eggs, 3 dl flour, 5 dl milk, butter')`)
	fmt.Println("Table insertion complete")
}

func serveHTTP() {
	mux := http.NewServeMux()

	log.Print("Listening on http://localhost:3000")
	//mux.HandleFunc("GET /recipe/all}", allRecipes)

	http.ListenAndServe(":3000", mux)
}

func allRecipes() {
	db, _ := sql.Open("sqlite3", "recipe.db")
	result, err := db.Query(`SELECT * FROM recipes`)
	if err != nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
