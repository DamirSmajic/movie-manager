package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	initDB()
	defer db.Close()

	userName := promptInput("Enter your name: ")

	currentUser, err := loadUser(userName)
	if err != nil {
		log.Fatalf("Error loading user: %v", err)
	}

	if currentUser == nil {
		fmt.Println("Creating a new user:", userName)
		currentUser = &User{Name: userName}
		Users[userName] = currentUser
	} else {
		fmt.Println("Welcome back,", userName)
	}

	for {
		fmt.Println("\nOptions: Print List (P), Add Movie (A), Change Rating (R), Save List (S), Exit (E)")
		choice := strings.ToUpper(promptInput("Choose an option: "))

		switch choice {
		case "P":
			printMovies(currentUser)
		case "A":
			addMovie(currentUser)
		case "R":
			changeRating(currentUser)
		case "S":
			saveList(currentUser)
			return
		case "E":
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
