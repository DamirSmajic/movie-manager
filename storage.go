package main

type Movie struct {
	Title    string
	Genre    string
	Rating   float32
	Director string
}

type MovieStorage interface {
	AddMovie(movie Movie) error
	GetMovies() ([]Movie, error)
}
