package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	colorBlue  = "\033[34m"
	colorReset = "\033[0m"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wordler",
	Short: "CLI Toolkit for the popular guessing game",
	Long: colorBlue + "                          ____\n _      ______  _________/ / /__  _____\n| | /| / / __ \\/ ___/ __  / / _ \\/ ___/\n| |/ |/ / /_/ / /  / /_/ / /  __/ /\n|__/|__/\\____/_/   \\__,_/_/\\___/_/\n\n" + colorReset +
		`The wordler CLI provides a number of helpful tools for the New York Times' hit game: Wordle.
Including the ability to play the game!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
