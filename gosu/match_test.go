package gosu

import (
	"fmt"
	"os"
)

func ExampleMatchCall() {
	mc := MatchCall{
		// Required:
		MatchID: "1936471",
	}

	fmt.Println(mc)
}

func ExampleMatch() {
	session := NewSession(os.Getenv("API-KEY"))

	call := MatchCall{
		MatchID: "1936471",
	}

	match, _ := session.FetchMatch(call)

	fmt.Println(match.Details.MatchID)
}
