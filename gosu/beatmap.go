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
	Approved string `json:"approved"`

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
	BPM string `json:"bpm"`

	// The creator of the beatmap.
	Creator string `json:"creator"`

	// ID of the beatmap's creator.
	CreatorID string `json:"creator_id"`

	// The star rating of the beatmap.
	DifficultyRating string `json:"difficultyrating"`

	// The circle size used in the beatmap.
	CircleSize string `json:"diff_size"`

	// The overall difficulty used in the beatmap.
	OverallDifficulty string `json:"diff_overall"`

	// The approach rate used in the beatmap.
	ApproachRate string `json:"diff_approach"`

	// The health drain used in the beatmap.
	HealthDrain string `json:"diff_drain"`

	// The number of seconds from the first note to the last note, not including breaks.
	HitLength string `json:"hit_length"`

	// The source of the song used in the beatmap.
	Source string `json:"source"`

	// ID of the genre of the song used in the beatmap.
	// 0 = any, 1 = unspecified, 2 = video game, 3 = anime, 4 = rock, 5 = pop, 6 = other, 7 = novelty, 9 = hip hop, 10 = electronic (note that there's no 8)
	GenreID string `json:"genre_id"`

	// ID of the language used in the song used in the beatmap.
	// 0 = any, 1 = other, 2 = english, 3 = japanese, 4 = chinese, 5 = instrumental, 6 = korean, 7 = french, 8 = german, 9 = swedish, 10 = spanish, 11 = italian
	LanguageID string `json:"language_id"`

	// The name of the song used in the beatmap.
	Title string `json:"title"`

	// The number of seconds from the first note to the last note, including breaks.
	TotalLength string `json:"total_length"`

	// The name of the beatmap's difficulty.
	Version string `json:"version"`

	// MD5 hash of the beatmap.
	FileMD5 string `json:"file_md5"`

	// The game mode the beatmap utilizes.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode string `json:"mode"`

	// The tags of the beatmap separated by spaces.
	Tags string `json:"tags"`

	// The number of times the beatmap has been favored.
	FavouriteCount string `json:"favourite_count"`

	// The number of times the beatmap has been played.
	PlayCount string `json:"playcount"`

	// The number of times the beatmap has been passed, completed.
	PassCount string `json:"passcount"`

	// The maximum combo a user can reach playing the beatmap.
	MaxCombo string `json:"max_combo"`
}

// FetchBeatmap returns metadata about one beatmap
func (s *Session) FetchBeatmap(call BeatmapCall) (Beatmap, error) {
	beatmap := new([]Beatmap)
	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)

	switch {
	case call.BeatmapID != "":
		v.Add(EndpointBeatmapsBeatmapID, call.BeatmapID)
	default:
		return Beatmap{}, errors.New("no identifying param given (BeatmapID)")
	}

	if call.Mode != "" {
		v.Add(EndpointBeatmapsMode, call.Mode)
	}
	if call.Converted != "" {
		v.Add(EndpointBeatmapsConverted, call.Converted)
	}
	if call.Hash != "" {
		v.Add(EndpointBeatmapsHash, call.Hash)
	}

	s.parseJSON(s.buildCall(EndpointBeatmaps, v), beatmap)

	if len(*beatmap) == 0 {
		return Beatmap{}, errors.New("no beatmaps found")
	}

	return (*beatmap)[0], nil
}
