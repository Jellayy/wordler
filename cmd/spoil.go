package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jellayy/wordler/utils"
	"github.com/spf13/cobra"
)

const (
	colorGreen = "\033[32m"
)

var (
	date string
)

var spoilCmd = &cobra.Command{
	Use:   "spoil",
	Short: "Spoil the answer to the NYT Wordle",
	Run:   spoil,
}

func init() {
	rootCmd.AddCommand(spoilCmd)
	spoilCmd.Flags().StringVarP(&date, "date", "d", "", "Date of Wordle to spoil, format: YYYY-MM-DD")
}

func spoil(cmd *cobra.Command, args []string) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	solution, err := utils.GrabNYTWord(date)
	if err != nil {
		fmt.Printf("Could not grab NYT solution for date: %s\n", date)
		return
	}

	fmt.Printf("%sWordle #%s - %s%s\n", colorBlue, strconv.Itoa(solution.Id), date, colorReset)
	fmt.Printf("Editor: %s\n", solution.Editor)
	fmt.Printf("Solution: %s%s%s\n", colorGreen, solution.Word, colorReset)
}
