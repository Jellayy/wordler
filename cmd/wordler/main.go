package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	colorGreen = "\033[32m"
	colorWhite = "\033[37m"
	colorBlue  = "\033[34m"
)

// Check if char in list
func charIn(a int32, list []int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Display game board
func display(board [][]int32, correct []int32) {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == correct[j] {
				fmt.Print(string(colorGreen) + string(board[i][j]) + string(colorWhite) + " ")
			} else if charIn(board[i][j], correct) {
				fmt.Print(string(colorBlue) + string(board[i][j]) + string(colorWhite) + " ")
			} else {
				fmt.Print(string(board[i][j]) + " ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	// Define wordle 5 letter words
	words := []string{
		"AUDIO",
		"SALET",
	}

	// Choose a random word
	word := words[rand.Intn(len(words))]

	// Create guessing arrays
	guesses := [][]int32{
		{'_', '_', '_', '_', '_'},
		{'_', '_', '_', '_', '_'},
		{'_', '_', '_', '_', '_'},
		{'_', '_', '_', '_', '_'},
		{'_', '_', '_', '_', '_'},
		{'_', '_', '_', '_', '_'},
	}
	correct := []int32{'_', '_', '_', '_', '_'}
	for i := 0; i < 5; i++ {
		correct[i] = int32(word[i])
	}

	display(guesses, correct)

	// Game loop
	victory := false
	for i := 0; i < 6; i++ {
		var guess string
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess)

		for j := 0; j < len(guess); j++ {
			guesses[i][j] = int32(guess[j])
		}

		display(guesses, correct)

		if guess == word {
			victory = true
			break
		}
	}

	if victory {
		fmt.Println("Congrats!")
	} else {
		fmt.Println("Unlucky\nThe word was: " + word)
	}
}
