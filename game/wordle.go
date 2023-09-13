package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jellayy/wordler/utils"
)

const (
	colorGreen = "\033[32m"
	colorWhite = "\033[37m"
	colorBlue  = "\033[34m"
	colorRed   = "\033[31m"
)

type game struct {
	gameMode          string
	prettyGameMode    string
	gameBoard         [][]int32
	answer            []int32
	AnswerStr         string
	numGuesses        int
	maxGuesses        int
	lastGuessFeedback string
	Victory           bool
	Completed         bool
}

func New(gameMode string) (game, error) {
	newGame := game{
		gameMode:       gameMode,
		prettyGameMode: "",
		gameBoard: [][]int32{
			{'_', '_', '_', '_', '_'},
			{'_', '_', '_', '_', '_'},
			{'_', '_', '_', '_', '_'},
			{'_', '_', '_', '_', '_'},
			{'_', '_', '_', '_', '_'},
			{'_', '_', '_', '_', '_'},
		},
		answer:            []int32{'_', '_', '_', '_', '_'},
		AnswerStr:         "",
		numGuesses:        0,
		maxGuesses:        6,
		lastGuessFeedback: "",
		Victory:           false,
		Completed:         false,
	}

	if gameMode == "nytdaily" {
		// nytdaily - NYT Daily Mode - Plays today's NYT wordle, validates with NYT dataset
		newGame.prettyGameMode = "NYT Daily Mode"
		newGame.AnswerStr = utils.GrabNYTWord()
	} else if gameMode == "random" {
		// random - Random Mode - Plays with random datamuse word, no dataset validation
		newGame.prettyGameMode = "Random Mode"
		newGame.AnswerStr = utils.GrabDatamuseWord()
	} else {
		return newGame, errors.New("unsupported gamemode")
	}

	newGame.answer = formatWord(newGame.AnswerStr)

	return newGame, nil
}

// -----------------
// Public functions
// -----------------

func (gameState game) DrawGame() {
	var gameDisplay string
	var spaceColor string

	// Build game display header with guesses and gamemode
	gameDisplay += colorWhite + "Guesses " + fmt.Sprint(gameState.numGuesses) + "/" + fmt.Sprint(gameState.maxGuesses) + " | " + gameState.prettyGameMode + "\n\n"

	// Build game board rows
	for row := 0; row < len(gameState.gameBoard); row++ {
		gameDisplay += "        "
		// Build row spaces
		for col := 0; col < len(gameState.gameBoard[row]); col++ {
			// Determine space color
			if gameState.gameBoard[row][col] == gameState.answer[col] {
				spaceColor = colorGreen
			} else if charIn(gameState.gameBoard[row][col], gameState.answer) {
				spaceColor = colorBlue
			} else {
				spaceColor = colorWhite
			}

			// Build space
			gameDisplay += spaceColor + string(gameState.gameBoard[row][col]) + colorWhite + " "
		}
		// End row
		gameDisplay += "\n"
	}

	// Clear & Build player input line or message
	if gameState.lastGuessFeedback != "" {
		gameDisplay += "\n\033[2K" + colorBlue + "> " + colorRed + gameState.lastGuessFeedback + colorWhite + ": "
	} else {
		gameDisplay += "\n\033[2K" + colorBlue + "> " + colorWhite + "Your Guess: "
	}

	// Move cursor to redraw over existing game display if already drawn
	if gameState.numGuesses > 0 {
		fmt.Printf("\033[%dA", countRune(gameDisplay, '\n')+1)
	}
	// Draw game display
	fmt.Print(gameDisplay)
}

func (gameState *game) ProcessGuess(guess string) {
	validGuess := true
	guess = strings.ToUpper(guess)

	// Length rule
	if len(guess) < 5 {
		gameState.lastGuessFeedback = "Not enough letters"
		validGuess = false
	} else if len(guess) > 5 {
		gameState.lastGuessFeedback = "Too many letters"
		validGuess = false
	}

	// Commit guess if valid
	if validGuess {
		for j := 0; j < len(guess); j++ {
			gameState.gameBoard[gameState.numGuesses][j] = int32(guess[j])
		}
		gameState.numGuesses++

		// End game if needed
		if guess == gameState.AnswerStr {
			gameState.Victory = true
			gameState.Completed = true
		} else if gameState.numGuesses >= gameState.maxGuesses {
			gameState.Completed = true
		}
	}
}

// --------------------------
// Internal helper functions
// --------------------------

// Split word into int32 array for game state
func formatWord(word string) []int32 {
	formattedWord := []int32{'_', '_', '_', '_', '_'}
	for i := 0; i < 5; i++ {
		formattedWord[i] = int32(word[i])
	}
	return formattedWord
}

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
