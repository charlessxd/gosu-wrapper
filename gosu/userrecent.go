package gosu

import (
	"errors"
	"net/url"
)

// UserRecentCall is used to build an API call to retrieve metadata on a user's recent plays.
type UserRecentCall struct {
	// ID or Username of the target user.
	UserID string

	// Specific game-mode.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode string

	// Maximum amount of results
	// Range between 1 - 50 (defaults to 10).
	Limit string

	// Whether UserID is an ID or a Username.
	// "id" if ID
	// "string" if username
	Type string
}

// UserRecent stores data for all recent plays an individual osu user has submitted.
type UserRecent struct {
	// ID of the beatmap.
	BeatmapID string `json:"beatmap_id"`

	// The score achieved on the beatmap.
	Score int64 `json:"score,string"`

	// The highest combo the user reached.
	MaxCombo int `json:"maxcombo,string"`

	// The number of 300s.
	Count300 int `json:"count300,string"`

	// The number of 100s.
	Count100 int `json:"count100,string"`

	// The number of 50s.
	Count50 int `json:"count50,string"`

	// The number of misses.
	CountMiss int `json:"countmiss,string"`

	// The number of katu.
	CountKatu int `json:"countkatu,string"`

	// The number of geki.
	CountGeki int `json:"countgeki,string"`

	// Whether the user achieved the maximum combo of the beatmap.
	// 1 = max combo achieved, 0 = max combo not achieved.
	Perfect string `json:"perfect"`

	// The bitwise flag representation of the mods used.
	EnabledMods string `json:"enabled_mods"`

	// The ID of the user.
	UserID string `json:"user_id"`

	// The date the top play was made, in UTC.
	Date string `json:"date"`

	// The letter ranking of the top play.
	Rank string `json:"rank"`
}

// FetchUserRecent returns metadata about a user's recent plays.
func (s *Session) FetchUserRecent(call UserRecentCall) ([]UserRecent, error) {
	userrecent := new([]UserRecent)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.UserID != "":
		v.Add(endpointParamUserID, call.UserID)
	default:
		return []UserRecent{}, errors.New("no identifying parameter given (UserID)")
	}

	if call.Mode != "" {
		v.Add(endpointParamMode, call.Mode)
	}
	if call.Type != "" {
		v.Add(endpointParamType, call.Type)
	}
	if call.Limit != "" {
		v.Add(endpointParamLimit, call.Limit)
	}

	err := s.parseJSON(s.buildCall(endpointUserRecent, v), userrecent)

	if err != nil {
		return *userrecent, err
	}
	if len(*userrecent) == 0 {
		return *userrecent, errors.New("user not found")
	}

	return *userrecent, nil
}
