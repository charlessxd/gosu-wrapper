package gosu

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleMatchCall() {
	_ := MatchCall{
		// Required:
		MatchID: "1936471",
	}
}

func ExampleMatch() {
	session := NewSession(os.Getenv("API-KEY"))

	call := MatchCall{
		MatchID: "1936471",
	}

	match, _ := session.FetchMatch(call)

	fmt.Println(json.MarshalIndent(match, "", "\t"))
	// Output:
	// {
	//	"match":{
	//		"match_id"     : "1936471",
	//		"name"         : "Marcin's game",
	//		"start_time"   : "2013-10-06 03:34:54", // in UTC
	//		"end_time"     : null             // null if not ended, date in UTC when match is disbanded
	//	},
	//	"games":[{
	//		"game_id"      : "45668898",
	//		"start_time"   : "2013-10-06 03:36:27", // in UTC
	//		"end_time"     : "2013-10-06 03:40:01", // in UTC
	//		"beatmap_id"   : "181717",
	//		"play_mode"    : "0",              // standard = 0, taiko = 1, ctb = 2, o!m = 3
	//		"match_type"   : "0",              // couldn't find
	//		"scoring_type" : "0",              // winning condition: score = 0, accuracy = 1, combo = 2, score v2 = 3
	//		"team_type"    : "0",              // Head to head = 0, Tag Co-op = 1, Team vs = 2, Tag Team vs = 3
	//		"mods"         : "0",              // global mods, see reference below
	//		"scores"       : [{
	//			"slot"          : "0",         // 0 based index of player's slot
	//			"team"          : "0",         // if mode doesn't support teams it is 0, otherwise 1 = blue, 2 = red
	//			"user_id"       : "722665",
	//			"score"         : "3415874",
	//			"maxcombo"      : "411",
	//			"rank"          : "0",         // not used
	//			"count50"       : "0",
	//			"count100"      : "11",
	//			"count300"      : "425",
	//			"countmiss"     : "1",
	//			"countgeki"     : "67",
	//			"countkatu"     : "9",
	//			"perfect"       : "0",        // full combo
	//			"pass"          : "1"         // if player failed at the end of the map it is 0, otherwise (pass or revive) it is 1
	//		},{ ... }  ...]                   // scores set in this game
	//	}, { ... },  ...]                     // games played in this match
	//}
}
