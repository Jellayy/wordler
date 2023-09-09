/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/jellayy/wordler/utils"
	"github.com/spf13/cobra"
)

var (
	randomMode bool
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play a game of Wordle in the command line",
	Run:   play,
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().BoolVarP(&randomMode, "random", "r", false, "Play with a random word instead of today's NYT word")
}

func play(cmd *cobra.Command, args []string) {
	// Set word
	var word string
	var gameMode string
	if randomMode {
		word = utils.GrabWord()
		gameMode = "Random Mode"
	} else {
		word = utils.GrabSolution()
		gameMode = "NYT Mode"
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

	utils.DrawGame(0, 6, gameMode, guesses, correct)

	// Game loop
	victory := false
	for i := 0; i < 6; i++ {
		var guess string
		fmt.Scan(&guess)
		guess = strings.ToUpper(guess)

		for j := 0; j < len(guess); j++ {
			guesses[i][j] = int32(guess[j])
		}

		utils.DrawGame(i+1, 6, gameMode, guesses, correct)

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
