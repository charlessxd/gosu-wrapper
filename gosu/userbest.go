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
	Mode string

	// Maximum amount of results
	// Range between 1 - 100 (defaults to 10).
	Limit string

	// Whether UserID is an ID or a Username.
	// "id" if ID
	// "string" if username
	Type string
}

// userBestPlay stores data for a defined amount of top plays for a specific osu user.
type userBestPlay struct {
	// The ID of the beatmap.
	BeatmapID string `json:"beatmap_id"`

	// The score achieved on the beatmap.
	Score int64 `json:"score,string"`

	// The highest combo the user achieved on the beatmap.
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

	// The nubmer of geki.
	CountGeki int `json:"countgeki,string"`

	// Whether the user achieved the maximum combo of the beatmap.
	// 1 = max combo achieved, 0 = max combo not achieved.
	Perfect string `json:"perfect"`

	// The bitwise flag representation of the mods used.
	ModsInt int64 `json:"enabled_mods,string"`
	EnabledMods []string

	// The ID of the user.
	UserID string `json:"user_id"`

	// The date the top play was made, in UTC.
	Date string `json:"date"`

	// The letter ranking of the top play.
	Rank string `json:"rank"`

	// PP rewarded for achieving the play, as a float value.
	PP float64 `json:"pp,string"`
}

// UserBest holds plays
type UserBest struct {
	Plays []userBestPlay

	// API Call URL.
	apiURL string

	// Session fetched from
	session *session

	apiCall UserBestCall
}

// FetchUserBest returns metadata about a user's highest rated plays.
func (s *session) FetchUserBest(call UserBestCall) (UserBest, error) {
	userbest := *new([]userBestPlay)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.UserID != "":
		v.Add(endpointParamUserID, call.UserID)
	default:
		return UserBest{}, errors.New("no identifying parameter given (UserID)")
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

	err := s.parseJSON(s.buildCall(endpointUserBest, v), &userbest)

	ub := *new(UserBest)
	ub.Plays = userbest

	for i := 0; i < len(ub.Plays); i++ {
		ub.Plays[i].EnabledMods = getMods(ub.Plays[i].ModsInt)
	}

	if err != nil {
		return ub, err
	}
	if len(userbest) == 0 {
		return ub, errors.New("user not found")
	}

	ub.apiURL = s.buildCall(endpointUserBest, v)
	ub.session = s

	return ub, nil
}

// Update updates a User's top plays.
func (ub *UserBest) Update() error {
	temp, err := ub.session.FetchUserBest(ub.apiCall)
	*ub = temp

	if err != nil {
		return err
	}
	return nil
}
