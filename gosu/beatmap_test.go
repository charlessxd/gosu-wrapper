package gosu

import (
	"fmt"
	"os"
)

func ExampleBeatmapCall() {
	bm := BeatmapCall{
		// Required:
		BeatmapID: "<ID of Beatmap>",

		// Optional:
		Mode:      "<Game-mode of Beatmap>",
		Converted: "<Converted>",
	}

	fmt.Println(bm)
}

func ExampleBeatmap() {
	session := NewSession(os.Getenv("API-KEY"))

	call := BeatmapCall{
		BeatmapID: "252002",
	}

	bm, _ := session.FetchBeatmap(call)

	fmt.Println(bm.BeatmapID)
}
