package main

import (
	"fmt"
)

type User struct {
	Name        string
	WatchedList []Movie
}

func printMovies(user *User) {
	if len(user.WatchedList) == 0 {
		fmt.Println("No movies in the watched list.")
		return
	}
	fmt.Println("Watched Movies:")
	for i, movie := range user.WatchedList {
		fmt.Printf("%d. Title: %s, Genre: %s, Rating: %.1f, Director: %s\n", i+1, movie.Title, movie.Genre, movie.Rating, movie.Director)
	}
}

func changeRating(user *User) {
	printMovies(user)
	if len(user.WatchedList) == 0 {
		return
	}

	indexInput := promptInput("Enter the number of the movie you want to change the rating for: ")
	var index int
	fmt.Sscanf(indexInput, "%d", &index)

	if index < 1 || index > len(user.WatchedList) {
		fmt.Println("Invalid movie number.")
		return
	}

	ratingInput := promptInput("Enter new rating: ")
	var newRating float32
	fmt.Sscanf(ratingInput, "%f", &newRating)

	user.WatchedList[index-1].Rating = newRating
	fmt.Println("Rating updated successfully.")
}

var Users = make(map[string]*User)
