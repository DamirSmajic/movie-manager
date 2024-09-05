package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(db *sql.DB) *SQLiteStorage {
	return &SQLiteStorage{db: db}
}

func (s *SQLiteStorage) AddMovie(movie Movie) error {
	_, err := s.db.Exec(
		"INSERT INTO movies (title, genre, rating, director) VALUES (?, ?, ?, ?)",
		movie.Title, movie.Genre, movie.Rating, movie.Director,
	)
	return err
}

func (s *SQLiteStorage) GetMovies() ([]Movie, error) {
	rows, err := s.db.Query("SELECT title, genre, rating, director FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.Title, &movie.Genre, &movie.Rating, &movie.Director); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
