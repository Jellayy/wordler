package utils

import (
	"context"
	"strings"
	"time"

	"github.com/carlmjohnson/requests"
)

type NYTRespone struct {
	Id       int    `json:"id"`
	Solution string `json:"solution"`
	Editor   string `json:"editor"`
}

func GrabNYTWord() string {
	// Query wordle solution from the NYT
	var response NYTRespone
	err := requests.
		URL("https://www.nytimes.com/svc/wordle/v2/" + time.Now().Format("2006-01-02") + ".json").
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		return "ERROR"
	}

	// Return just the word for now
	// TODO: Return more data about the solution
	return strings.ToUpper(response.Solution)
}
