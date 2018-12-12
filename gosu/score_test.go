package gosu

import (
	"fmt"
	"os"
)

func ExampleScore() {
	session := NewSession(os.Getenv("API_KEY"))

	c := ScoreCall{
		BeatmapID: os.Getenv("BEATMAP_ID"),
	}

	scores, _ := session.FetchScores(c)

	if len(scores) > 0 {
		fmt.Println(scores[0].UserID)
	}
}
