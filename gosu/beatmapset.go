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

type Beatmaps struct {
	Beatmaps []Beatmap

	// Whether converted beatmaps are included
	// 0 = not included, 1 = included
	Converted string

	// API Call URL.
	apiURL string

	// Session fetched from
	session *session

	apiCall BeatmapsCall
}

// FetchBeatmaps returns metadata about multiple beatmaps.
func (s *session) FetchBeatmaps(call BeatmapsCall) (Beatmaps, error) {
	beatmaps := *new([]Beatmap)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.BeatmapSetID != "":
		v.Add(endpointParamBeatmapSetID, call.BeatmapSetID)
	case call.Since != "":
		v.Add(endpointParamSince, call.Since)
	default:
		return Beatmaps{}, errors.New("no identifying param given (Since, BeatmapSetID)")
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
	if call.Limit != "" {
		v.Add(endpointParamLimit, call.Limit)
	}
	if call.Type != "" {
		v.Add(endpointParamType, call.Type)
	}

	err := s.parseJSON(s.buildCall(endpointBeatmaps, v), &beatmaps)

	if err != nil {
		return Beatmaps{}, err
	}
	if len(beatmaps) == 0 {
		return Beatmaps{}, errors.New("no beatmaps found")
	}

	set := *new(Beatmaps)

	set.Beatmaps = beatmaps
	set.apiURL = s.buildCall(endpointBeatmaps, v)
	set.session = s
	set.Converted = call.Converted

	// Allows for each beatmap in the set to be updated individually.
	for i := 0; i < len(set.Beatmaps); i++ {
		v = url.Values{}

		if set.Converted != "" {
			v.Add(endpointParamConverted, call.Converted)
		}

		v.Add(endpointParamBeatmapID, set.Beatmaps[i].BeatmapID)
		v.Add(endpointParamMode, set.Beatmaps[i].Mode)
		set.Beatmaps[i].apiURL = set.apiURL
		set.Beatmaps[i].session = set.session
	}

	return set, nil
}

// Update updates a Beatmapset
func (set *Beatmaps) Update() error {
	temp, err := set.session.FetchBeatmaps(set.apiCall)
	*set = temp

	if err != nil {
		return err
	}
	return nil

	// Allows for the updating of individual beatmaps
	for i := 0; i < len(set.Beatmaps); i++ {
		v := url.Values{}
		v.Add(endpointAPIKey, set.session.key)

		if set.Converted != "" {
			v.Add(endpointParamConverted, set.Converted)
		}

		v.Add(endpointParamBeatmapID, set.Beatmaps[i].BeatmapID)
		v.Add(endpointParamMode, set.Beatmaps[i].Mode)
		set.Beatmaps[i].apiURL = set.session.buildCall(endpointBeatmaps, v)
		set.Beatmaps[i].session = set.session
	}

	return nil
}
