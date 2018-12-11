package gosu

import (
	"errors"
	"net/url"
)

// UserBestCall is used to build an API call to retrieve metadata on a user's top plays.
type UserBestCall struct {
	// ID or Username of the target user.
	UserID string

	// Specific game-mode.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode   string

	// Maximum amount of results
	// Range between 1 - 100 (defaults to 10).
	Limit  string

	// Whether UserID is an ID or a Username.
	// "id" if ID
	// "string" if username
	Type   string
}

// UserBest stores data for a defined amount of top plays for a specific osu user.
type UserBest struct {
	// The ID of the beatmap.
	BeatmapID string `json:"beatmap_id"`

	// The score achieved on the beatmap.
	Score string `json:"score"`

	// The highest combo the user achieved on the beatmap.
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

	// The nubmer of geki.
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

	// PP rewarded for achieving the play, as a float value.
	PP string `json:"pp"`
}

// FetchUserBest returns metadata about a user's highest rated plays.
func (s *Session) FetchUserBest(call UserBestCall) ([]UserBest, error) {
	userbest := new([]UserBest)
	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)

	switch {
	case call.UserID != "":
		v.Add(EndpointUserBestUserID, call.UserID)
	default:
		return []UserBest{}, errors.New("no identifying parameter given (UserID)")
	}

	if call.Mode != "" {
		v.Add(EndpointUserBestMode, call.Mode)
	}
	if call.Type != "" {
		v.Add(EndpointUserBestType, call.Type)
	}
	if call.Limit != "" {
		v.Add(EndpointUserBestLimit, call.Limit)
	}

	s.parseJSON(s.buildCall(EndpointUserBest, v), userbest)

	if len(*userbest) == 0 {
		return *userbest, errors.New("user not found")
	}

	return *userbest, nil
}
