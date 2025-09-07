package main

import (
	"aprende-golang/internal/service"
	"aprende-golang/internal/store"
	"aprende-golang/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to SQLite
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist

	q := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	)`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inject our dependencies

	bookStore := store.New(db)
	bookService := service.New(bookStore)
	BookHandler := transport.New(bookService)

	// Configure the HTTP server

	http.HandleFunc("/books", BookHandler.HandleBooks)
	http.HandleFunc("/books/", BookHandler.HandleBookByID)

	fmt.Println("Server is running on port 8000")

	// Start and listen the server
	log.Fatal(http.ListenAndServe(":8000", nil))

}
