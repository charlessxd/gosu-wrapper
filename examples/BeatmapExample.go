package main

import (
	"../gosu"
	"fmt"
)

func main() {
	s := gosu.NewSession("<API Key>")

	call := gosu.BeatmapCall{
		BeatmapID: "<BeatmapID>",
	}

	beatmap, _ := s.FetchBeatmap(call)

	fmt.Println(beatmap)
}
