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
type UserRecent []struct {
	// ID of the beatmap.
	BeatmapID string `json:"beatmap_id"`

	// The score achieved on the beatmap.
	Score string `json:"score"`

	// The highest combo the user reached.
	MaxCombo string `json:"maxcombo"`

	// The number of 300s.
	Count300 string `json:"count300"`

	// The number of 100s.
	Count100 string `json:"count100"`

	// The number of 50s.
	Count50 string `json:"count50"`

	// The number of misses.
	CountMiss string `json:"countmiss"`

	// The number of katu.
	CountKatu string `json:"countkatu"`

	// The number of geki.
	CountGeki string `json:"countgeki"`

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
func (s *Session) FetchUserRecent(call UserRecentCall) (UserRecent, error) {
	userrecent := new(UserRecent)
	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)

	switch {
	case call.UserID != "":
		v.Add(EndpointUserRecentUserID, call.UserID)
	default:
		return UserRecent{}, errors.New("no identifying parameter given (UserID)")
	}

	if call.Mode != "" {
		v.Add(EndpointUserRecentMode, call.Mode)
	}
	if call.Type != "" {
		v.Add(EndpointUserRecentType, call.Type)
	}
	if call.Limit != "" {
		v.Add(EndpointUserRecentLimit, call.Limit)
	}

	s.ParseJSON(s.BuildCall(EndpointUserRecent, v), userrecent)

	if len(*userrecent) == 0 {
		return *userrecent, errors.New("user not found")
	}

	return *userrecent, nil
}
