package gosu

import (
	"errors"
	"net/url"
)

// BeatmapCall is used to build an API call to retrieve metadata on one beatmap.
type BeatmapCall struct {
	// ID of the beatmap
	BeatmapID string

	// Specific game-mode.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode string

	// Whether converted beatmaps are included
	// 0 = not included, 1 = included
	Converted string

	// The beatmap hash
	Hash string
}

// Beatmap stores the data of a beatmap.
type Beatmap struct {
	// The status of the beatmap's ranking.
	// 4 = loved, 3 = qualified, 2 = approved, 1 = ranked, 0 = pending, -1 = WIP, -2 = graveyard
	Approved int `json:"approved,string"`

	// Date the beatmap was ranked, in UTC.
	ApprovedDate string `json:"approved_date"`

	// Date the beatmap was last updated, in UTC.
	LastUpdate string `json:"last_update"`

	// Artist of the song used in the beatmap.
	Artist string `json:"artist"`

	// ID of the beatmap.
	BeatmapID string `json:"beatmap_id"`

	// ID of the beatmap set the beatmap is contained in.
	BeatmapSetID string `json:"beatmapset_id"`

	// The BPM of the beatmap.
	BPM int `json:"bpm,string"`

	// The creator of the beatmap.
	Creator string `json:"creator"`

	// ID of the beatmap's creator.
	CreatorID string `json:"creator_id"`

	// The star rating of the beatmap.
	DifficultyRating float64 `json:"difficultyrating,string"`

	// The circle size used in the beatmap.
	CircleSize float64 `json:"diff_size,string"`

	// The overall difficulty used in the beatmap.
	OverallDifficulty float64 `json:"diff_overall,string"`

	// The approach rate used in the beatmap.
	ApproachRate float64 `json:"diff_approach,string"`

	// The health drain used in the beatmap.
	HealthDrain float64 `json:"diff_drain,string"`

	// The number of seconds from the first note to the last note, not including breaks.
	HitLength int `json:"hit_length,string"`

	// The source of the song used in the beatmap.
	Source string `json:"source"`

	// ID of the genre of the song used in the beatmap.
	// 0 = any, 1 = unspecified, 2 = video game, 3 = anime, 4 = rock, 5 = pop, 6 = other, 7 = novelty, 9 = hip hop, 10 = electronic (note that there's no 8)
	GenreID int `json:"genre_id,string"`

	// ID of the language used in the song used in the beatmap.
	// 0 = any, 1 = other, 2 = english, 3 = japanese, 4 = chinese, 5 = instrumental, 6 = korean, 7 = french, 8 = german, 9 = swedish, 10 = spanish, 11 = italian
	LanguageID int `json:"language_id,string"`

	// The name of the song used in the beatmap.
	Title string `json:"title"`

	// The number of seconds from the first note to the last note, including breaks.
	TotalLength int `json:"total_length,string"`

	// The name of the beatmap's difficulty.
	Version string `json:"version"`

	// MD5 hash of the beatmap.
	FileMD5 string `json:"file_md5"`

	// The game mode the beatmap utilizes.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode int `json:"mode,string"`

	// The tags of the beatmap separated by spaces.
	Tags string `json:"tags"`

	// The number of times the beatmap has been favored.
	FavouriteCount int `json:"favourite_count,string"`

	// The number of times the beatmap has been played.
	PlayCount int `json:"playcount,string"`

	// The number of times the beatmap has been passed, completed.
	PassCount int `json:"passcount,string"`

	// The maximum combo a user can reach playing the beatmap.
	MaxCombo int `json:"max_combo,string"`
}

// FetchBeatmap returns metadata about one beatmap
func (s *Session) FetchBeatmap(call BeatmapCall) (Beatmap, error) {
	beatmap := new([]Beatmap)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.BeatmapID != "":
		v.Add(endpointParamBeatmapID, call.BeatmapID)
	default:
		return Beatmap{}, errors.New("no identifying param given (BeatmapID)")
	}

	if call.Mode != "" {
		v.Add(endpointParamMode, call.Mode)
	}
	if call.Converted != "" {
		v.Add(endpointParamConverted, call.Converted)
	}
	if call.Hash != "" {
		v.Add(endpointParamHash, call.Hash)
	}

	err := s.parseJSON(s.buildCall(endpointBeatmaps, v), beatmap)

	if err != nil {
		return Beatmap{}, err
	}
	if len(*beatmap) == 0 {
		return Beatmap{}, errors.New("no beatmaps found")
	}

	return (*beatmap)[0], nil
}
