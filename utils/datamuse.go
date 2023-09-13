package utils

import (
	"context"
	"math/rand"
	"strings"

	"github.com/carlmjohnson/requests"
)

type DatamuseResponse struct {
	Score int    `json:"score"`
	Word  string `json:"word"`
}

func GrabDatamuseWord() string {
	// Query words from Datamuse
	var response []DatamuseResponse
	err := requests.
		URL("https://api.datamuse.com/words?sp=?????&max=500").
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		return "ERROR"
	}

	// Choose and return random word
	return strings.ToUpper(response[rand.Intn(len(response))].Word)
}
