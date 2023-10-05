package utils

import (
	"context"
	"strings"

	"github.com/carlmjohnson/requests"
)

type NYTRespone struct {
	Id     int    `json:"id"`
	Word   string `json:"solution"`
	Editor string `json:"editor"`
}

func GrabNYTWord(date string) (NYTRespone, error) {
	// Query wordle solution from the NYT
	var response NYTRespone
	err := requests.
		URL("https://www.nytimes.com/svc/wordle/v2/" + date + ".json").
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		return response, err
	}

	response.Word = strings.ToUpper(response.Word)

	return response, nil
}
