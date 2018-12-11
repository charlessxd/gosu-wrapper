package main

import (
	"../gosu"
	"fmt"
)

func main() {
	s := gosu.NewSession("18882e259698eeedc2e5ee310db1303380761d70")

	call := gosu.BeatmapCall{
		BeatmapID: "252002",
	}

	beatmap, _ := s.FetchBeatmap(call)

	fmt.Println(beatmap)
}
