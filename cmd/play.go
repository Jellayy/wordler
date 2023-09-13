package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jellayy/wordler/game"
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

	// Create game
	gameState, err := game.New(settings.GameMode)
	if err != nil {
		print(err)
		return
	}

	// Run game
	gameState.DrawGame()
	for !gameState.Completed {
		var guess string
		fmt.Scan(&guess)
		gameState.ProcessGuess(guess)
		gameState.DrawGame()
	}

	// Finish game
	if gameState.Victory {
		fmt.Println("\nCongrats!")
	} else {
		fmt.Println("\nUnlucky\nThe word was: " + gameState.AnswerStr)
	}
}
