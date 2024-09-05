package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestAddMovie(t *testing.T) {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE movies (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            genre TEXT,
            rating REAL,
            director TEXT
        );
    `)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	storage := NewSQLiteStorage(db)
	movie := Movie{Title: "Inception", Genre: "Sci-Fi", Rating: 8.8, Director: "Christopher Nolan"}

	err = storage.AddMovie(movie)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	movies, err := storage.GetMovies()
	if err != nil {
		t.Errorf("Failed to get movies: %v", err)
	}

	if len(movies) != 1 {
		t.Errorf("Expected 1 movie, got %d", len(movies))
	}

	if movies[0] != movie {
		t.Errorf("Expected %v, got %v", movie, movies[0])
	}
}
