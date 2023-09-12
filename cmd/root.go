/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wordler.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
