package gosu

import (
	"errors"
	"net/url"
)

// BeatmapsCall is used to build an API call to retrieve metadata on multiple beatmaps
type BeatmapsCall struct {
	// Return all beatmaps ranked or loved since this date.
	// MySQL date in UTC.
	Since string

	// ID of the target Beatmap Set
	BeatmapSetID string

	// ID or Username of the target user.
	UserID string

	// Specific game-mode.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode string

	// Whether converted beatmaps are included
	// 0 = not included, 1 = included
	Converted string

	// Amount of results
	// default and max is 500
	Limit string

	// Whether UserID is an ID or a Username.
	// "id" if ID
	// "string" if username
	Type string
}

// FetchBeatmaps returns metadata about multiple beatmaps.
func (s *Session) FetchBeatmaps(call BeatmapsCall) ([]Beatmap, error) {
	beatmaps := new([]Beatmap)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.BeatmapSetID != "":
		v.Add(endpointParamBeatmapSetID, call.BeatmapSetID)
	case call.Since != "":
		v.Add(endpointParamSince, call.Since)
	default:
		return *beatmaps, errors.New("no identifying param given (Since, BeatmapSetID)")
	}

	if call.UserID != "" {
		v.Add(endpointParamUserID, call.UserID)
	}
	if call.Mode != "" {
		v.Add(endpointParamMode, call.Mode)
	}
	if call.Converted != "" {
		v.Add(endpointParamConverted, call.Converted)
	}
	if call.Limit != "" && call.BeatmapSetID != "" {
		v.Add(endpointParamLimit, call.Limit)
	}
	if call.Type != "" {
		v.Add(endpointParamType, call.Type)
	}

	s.parseJSON(s.buildCall(endpointBeatmaps, v), beatmaps)

	if len(*beatmaps) == 0 {
		return *beatmaps, errors.New("no beatmaps found")
	}

	return *beatmaps, nil
}
