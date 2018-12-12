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
	session := NewSession(os.Getenv("API_KEY"))

	call := BeatmapCall{
		BeatmapID: "BEATMAP_ID	",
	}

	bm, _ := session.FetchBeatmap(call)

	fmt.Println(bm.BeatmapID)
}
