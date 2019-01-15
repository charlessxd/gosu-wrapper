package gosu

import (
	"fmt"
	"os"
)

func ExampleSession_FetchScores() {
	session := NewSession(os.Getenv("API_KEY"))

	c := ScoresCall{
		BeatmapID: os.Getenv("BEATMAP_ID"),
	}

	scores, _ := session.FetchScores(c)

	if len(scores.Scores) > 0 {
		fmt.Println(scores.Scores[0].UserID)
	}
}
