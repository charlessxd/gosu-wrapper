package gosu

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleBeatmapCall() {
	_ := BeatmapCall{
		// Required:
		BeatmapID: "<ID of Beatmap>",

		// Optional:
		Mode:      "<Game-mode of Beatmap>",
		Converted: "<Converted>",
	}
}

func ExampleBeatmap() {
	session := NewSession(os.Getenv("API-KEY"))

	call := BeatmapCall{
		BeatmapID: "252002",
	}

	bm, _ := session.FetchBeatmap(call)

	fmt.Println(json.MarshalIndent(bm, "", "\t"))
	// Output:
	// {
	//	"approved"         : "1",                   // 4 = loved, 3 = qualified, 2 = approved, 1 = ranked, 0 = pending, -1 = WIP, -2 = graveyard
	//	"approved_date"    : "2013-07-02 01:01:12", // date ranked, in UTC
	//	"last_update"      : "2013-07-06 16:51:22", // last update date, in UTC. May be after approved_date if map was unranked and reranked.
	//	"artist"           : "Luxion",
	//	"beatmap_id"       : "252002",              // beatmap_id is per difficulty
	//	"beatmapset_id"    : "93398",               // beatmapset_id groups difficulties into a set
	//	"bpm"              : "196",
	//	"creator"          : "RikiH_",
	//	"creator_id"       : "686209",
	//	"difficultyrating" : "5.59516",             // The amount of stars the map would have ingame and on the website
	//	"diff_size"        : "4",                   // Circle size value (CS)
	//	"diff_overall"     : "6",                   // Overall difficulty (OD)
	//	"diff_approach"    : "7",                   // Approach Rate (AR)
	//	"diff_drain"       : "6",                   // Healthdrain (HP)
	//	"hit_length"       : "113",                 // seconds from first note to last note not including breaks
	//	"source"           : "BMS",
	//	"genre_id"         : "1",                   // 0 = any, 1 = unspecified, 2 = video game, 3 = anime, 4 = rock, 5 = pop, 6 = other, 7 = novelty, 9 = hip hop, 10 = electronic (note that there's no 8)
	//	"language_id"      : "5",                   // 0 = any, 1 = other, 2 = english, 3 = japanese, 4 = chinese, 5 = instrumental, 6 = korean, 7 = french, 8 = german, 9 = swedish, 10 = spanish, 11 = italian
	//	"title"            : "High-Priestess",      // song name
	//	"total_length"     : "145",                 // seconds from first note to last note including breaks
	//	"version"          : "Overkill",            // difficulty name
	//	"file_md5"         : "c8f08438204abfcdd1a748ebfae67421", // md5 hash of the beatmap
	//	"mode"             : "0",                   // game mode,
	//	"tags"             : "melodious long",      // Beatmap tags separated by spaces.
	//	"favourite_count"  : "121",                 // Number of times the beatmap was favourited. (americans: notice the ou!)
	//	"playcount"        : "9001",                // Number of times the beatmap was played
	//	"passcount"        : "1337",                // Number of times the beatmap was passed, completed (the user didn't fail or retry)
	//	"max_combo"        : "2101"                 // The maximum combo a user can reach playing this beatmap.
	//}
}
