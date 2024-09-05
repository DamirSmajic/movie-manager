package main

import "fmt"

func addMovie(user *User) {
	title := promptInput("Enter movie title: ")
	genre := promptInput("Enter movie genre: ")
	ratingInput := promptInput("Enter movie rating: ")
	var rating float32
	fmt.Sscanf(ratingInput, "%f", &rating)
	director := promptInput("Enter movie director: ")

	movie := Movie{Title: title, Genre: genre, Rating: rating, Director: director}
	user.WatchedList = append(user.WatchedList, movie)
	fmt.Println("Movie added successfully.")
}
