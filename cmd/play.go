/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jellayy/wordler/utils"
	"github.com/spf13/cobra"
)

var (
	settings struct {
		GameMode string
	}
	selectMode = []*survey.Question{
		{
			Name: "gameMode",
			Prompt: &survey.Select{
				Message: "Choose a Game Mode:",
				Options: []string{"nytdaily", "random"},
				Description: func(value string, index int) string {
					if value == "nytdaily" {
						return "Today's New York Times daily Wordle"
					} else if value == "random" {
						return "Random word from the datamuse dataset"
					}
					return ""
				},
			},
		},
	}
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a game of Wordle in the command line",
	Run:   play,
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVarP(&settings.GameMode, "gamemode", "g", "", "Choose gamemode, available options: nytdaily, random")
}

func play(cmd *cobra.Command, args []string) {
	// Query for gamemode if not set in args
	if settings.GameMode == "" {
		err := survey.Ask(selectMode, &settings)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// Set word & gamemode
	var word string
	var prettyGameMode string
	if settings.GameMode == "nytdaily" {
		word = utils.GrabSolution()
		prettyGameMode = "NYT Daily Mode"
	} else if settings.GameMode == "random" {
		word = utils.GrabWord()
		prettyGameMode = "Random Mode"
	} else {
		fmt.Print("ERROR: Gamemode not supported")
	}

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

	utils.DrawGame(0, 6, prettyGameMode, guesses, correct, "")

	// Game loop
	victory := false
	numGuesses := 0
	for numGuesses < 6 {
		var guess string
		message := ""
		validGuess := true

		// Read guess
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess)

		// Validate guess
		if len(guess) < 5 {
			message = "Not enough letters"
			validGuess = false
		} else if len(guess) > 5 {
			message = "Too many letters"
			validGuess = false
		}

		// Commit guess if valid
		if validGuess {
			for j := 0; j < len(guess); j++ {
				guesses[numGuesses][j] = int32(guess[j])
			}
			numGuesses++
		}

		utils.DrawGame(numGuesses, 6, prettyGameMode, guesses, correct, message)

		if guess == word {
			victory = true
			break
		}
	}

	if victory {
		fmt.Println("\nCongrats!")
	} else {
		fmt.Println("\nUnlucky\nThe word was: " + word)
	}
}
