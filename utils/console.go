package utils

import (
	"fmt"
)

const (
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

// Return count of rune in string, ex: \n
func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}

func DrawGame(guesses int, maxGuesses int, gameMode string, gameBoard [][]int32, solution []int32) {
	var gameDisplay string
	var spaceColor string

	// Build game display header with guesses and gamemode
	gameDisplay += "\nGuesses " + fmt.Sprint(guesses) + "/" + fmt.Sprint(maxGuesses) + " | " + gameMode + "\n\n"

	// Build game board rows
	for row := 0; row < len(gameBoard); row++ {
		gameDisplay += "        "
		// Build row spaces
		for col := 0; col < len(gameBoard[row]); col++ {
			// Determine space color
			if gameBoard[row][col] == solution[col] {
				spaceColor = colorGreen
			} else if charIn(gameBoard[row][col], solution) {
				spaceColor = colorBlue
			} else {
				spaceColor = colorWhite
			}

			// Build space
			gameDisplay += spaceColor + string(gameBoard[row][col]) + colorWhite + " "
		}
		// End row
		gameDisplay += "\n"
	}

	// Clear & Build player input line
	gameDisplay += "\n\033[2K" + colorBlue + ">" + colorWhite + " Your Guess: "

	// Move cursor to redraw over existing game display if already drawn
	if guesses > 0 {
		fmt.Printf("\033[%dA", countRune(gameDisplay, '\n')+1)
	}
	// Draw game display
	fmt.Print(gameDisplay)
}
