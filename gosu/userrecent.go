package gosu

import (
	"errors"
	"net/url"
	"strconv"
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

// userRecentPlay stores data for all recent plays an individual osu user has submitted.
type UserRecentPlay struct {
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
	ModsInt int64 `json:"enabled_mods,string"`
	EnabledMods []string

	// The ID of the user.
	UserID string `json:"user_id"`

	// The date the top play was made, in UTC.
	Date string `json:"date"`

	// The letter ranking of the top play.
	Rank string `json:"rank"`

	Accuracy float64
}

// UserRecent stores UserRecentPlays and allows for them to be updated using Update.
type UserRecent struct {
	Plays []UserRecentPlay

	// API Call URL.
	apiURL string

	// Session fetched from
	session *session

	apiCall UserRecentCall
}

// FetchUserRecent returns metadata about a user's recent plays.
func (s *session) FetchUserRecent(call UserRecentCall) (UserRecent, error) {
	if i, e := strconv.ParseInt(call.Limit, 10, 64); i > 100 {
		return UserRecent{}, errors.New("limit parameter exceeds maximum limit (Max limit: 100)")
	} else if e != nil {
		return UserRecent{}, e
	}

	plays := *new([]UserRecentPlay)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.UserID != "":
		v.Add(endpointParamUserID, call.UserID)
	default:
		return UserRecent{}, errors.New("no identifying parameter given (UserID)")
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

	err := s.parseJSON(s.buildCall(endpointUserRecent, v), &plays)

	userrecent := *new(UserRecent)
	userrecent.Plays = plays

	if err != nil {
		return userrecent, err
	}
	if len(userrecent.Plays) == 0 {
		return userrecent, errors.New("user not found")
	}

	for i, play := range userrecent.Plays {
		userrecent.Plays[i].EnabledMods = getMods(play.ModsInt)
		userrecent.Plays[i].Accuracy = getAcc(play.CountMiss, play.Count50, play.Count100, play.Count300)
	}

	userrecent.apiURL = s.buildCall(endpointUserRecent, v)
	userrecent.session = s

	return userrecent, nil
}

// Update updates a User's recent plays.
func (ur *UserRecent) Update() error {
	temp, err := ur.session.FetchUserRecent(ur.apiCall)
	*ur = temp

	if err != nil {
		return err
	}
	return nil
}
