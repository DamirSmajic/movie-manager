package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./movies.db")
	if err != nil {
		log.Fatal(err)
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE
	);
	`

	createMovieTable := `
	CREATE TABLE IF NOT EXISTS movies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		genre TEXT,
		rating REAL,
		director TEXT,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`

	_, err = db.Exec(createUserTable)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	_, err = db.Exec(createMovieTable)
	if err != nil {
		log.Fatalf("Error creating movies table: %v", err)
	}
}

func loadUser(name string) (*User, error) {

	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE name = ?", name).Scan(&userID)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to query user: %v", err)
	}

	rows, err := db.Query("SELECT title, genre, rating, director FROM movies WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to load movies: %v", err)
	}
	defer rows.Close()

	var watchedList []Movie
	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.Title, &movie.Genre, &movie.Rating, &movie.Director)
		if err != nil {
			return nil, fmt.Errorf("failed to scan movie: %v", err)
		}
		watchedList = append(watchedList, movie)
	}

	return &User{Name: name, WatchedList: watchedList}, nil
}

func saveList(user *User) {

	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE name = ?", user.Name).Scan(&userID)

	if err == sql.ErrNoRows {
		result, err := db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
		if err != nil {
			log.Fatalf("Failed to insert user: %v", err)
		}
		userID64, _ := result.LastInsertId()
		userID = int(userID64)
	} else if err != nil {
		log.Fatalf("Failed to query user: %v", err)
	}

	for _, movie := range user.WatchedList {
		_, err = db.Exec(
			"INSERT INTO movies (title, genre, rating, director, user_id) VALUES (?, ?, ?, ?, ?)",
			movie.Title, movie.Genre, movie.Rating, movie.Director, userID,
		)
		if err != nil {
			log.Fatalf("Failed to insert movie: %v", err)
		}
	}

	fmt.Println("List saved successfully.")
}
