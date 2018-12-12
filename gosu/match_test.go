package gosu

import (
	"fmt"
	"os"
)

func ExampleMatchCall() {
	mc := MatchCall{
		// Required:
		MatchID: "MATCH_ID",
	}

	fmt.Println(mc)
}

func ExampleMatch() {
	session := NewSession(os.Getenv("API_KEY"))

	call := MatchCall{
		MatchID: "MATCH_ID",
	}

	match, _ := session.FetchMatch(call)

	fmt.Println(match.Details.MatchID)
}
