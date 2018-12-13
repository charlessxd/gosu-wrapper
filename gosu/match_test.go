package gosu

import (
	"fmt"
	"os"
)

func ExampleSession_FetchMatch() {
	session := NewSession(os.Getenv("API_KEY"))

	c := MatchCall{
		MatchID: os.Getenv("MATCH_ID"),
	}

	match, _ := session.FetchMatch(c)

	fmt.Println(match.Details.MatchID)
}
