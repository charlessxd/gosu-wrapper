package gosu

import (
	"errors"
	"net/url"
)

// UserCall is used to build an API call to retrieve metadata of a user.
type UserCall struct {
	// ID of the target user.
	UserID string

	// Specific game-mode.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode string

	// Whether UserID is an ID or a Username.
	// "id" if ID
	// "string" if username
	Type string
}

// User stores data for an individual osu user.
type User struct {
	// ID of the user.
	UserID string `json:"user_id"`

	// The name of the user.
	Username string `json:"username"`

	// The number of 300s the user has gotten for all ranked, approved, and loved beatmaps played.
	Count300 int `json:"count300,string"`

	// The number of 100s the user has gotten for all ranked, approved, and loved beatmaps played.
	Count100 int `json:"count100,string"`

	// The number of 50s the user has gotten for all ranked, approved, and loved beatmaps played.
	Count50 int `json:"count50,string"`

	// The amount of plays the user has on ranked, approved, and loved beatmaps.
	PlayCount int `json:"playcount,string"`

	// The total of the best individual score for every ranked, approved, and loved beatmap the user has played.
	RankedScore int64 `json:"ranked_score,string"`

	// The total score of every beatmap the user has played.
	TotalScore int64 `json:"total_score,string"`

	// The global ranking of the user in terms of PP.
	// 1 having the highest amount of PP.
	PPRank int64 `json:"pp_rank,string"`

	// The level of the user.
	Level float64 `json:"level,string"`

	// The amount of PP the user has as a float value.
	PPRaw float64 `json:"pp_raw,string"`

	// The accuracy of the user as a percentage.
	Accuracy float64 `json:"accuracy,string"`

	// The total number of non-hidden SS ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankSS int `json:"count_rank_ss,string"`

	// The total number of hidden SS ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankSSH int `json:"count_rank_ssh,string"`

	// The total number of non-hidden S ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankS int `json:"count_rank_s,string"`

	// The total number of hidden S ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankSH int `json:"count_rank_sh,string"`

	// The total number of A ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankA int `json:"count_rank_a,string"`

	// The country the user created their account from.
	// Uses the ISO3166-1 alpha-2 country code naming. See this for more information: http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2/)
	Country string `json:"country"`

	// The total number of seconds the user has played.
	TotalSecondsPlayed int64 `json:"total_seconds_played,string"`

	// The PP ranking of the user within their country.
	PPCountryRank int `json:"pp_country_rank,string"`

	// The events of the user.
	Events []OsuEvent `json:"events"`

	// API Call URL
	apiURL string

	session *Session

	listeners []UserHandler
}

// UserEvent stores data for events related to an individual osu user.
type OsuEvent struct {
	// The HTML for the event.
	DisplayHTML string `json:"display_html"`

	// The ID of the beatmap in the event.
	BeatmapID string `json:"beatmap_id"`

	// The ID of beatmap set in the event.
	BeatmapSetID string `json:"beatmapset_id"`

	// The date the event occurred, in UTC.
	Date string `json:"date"`

	// How epic the event is, between 1 and 32.
	EpicFactor int `json:"epicfactor,string"`
}

type UserEvent struct {
}

// FetchUser returns metadata about a user.
func (s *Session) FetchUser(call UserCall) (User, error) {
	user := *new([]User)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.UserID != "":
		v.Add(endpointParamUserID, call.UserID)
	default:
		return User{}, errors.New("no identifying parameter given (UserID)")
	}

	if call.Mode != "" {
		v.Add(endpointParamMode, call.Mode)
	}
	if call.Type != "" {
		v.Add(endpointParamType, call.Type)
	}

	err := s.parseJSON(s.buildCall(endpointUser, v), &user)

	if err != nil {
		return User{}, err
	}
	if len(user) == 0 {
		return User{}, errors.New("user not found")
	}

	user[0].apiURL = s.buildCall(endpointUser, v)
	user[0].session = s

	return user[0], nil
}

func (u *User) Update() error {
	user := *new([]User)

	err := u.session.parseJSON(u.apiURL, &user)

	if err != nil {
		return err
	}
	if u.apiURL == "" {
		return errors.New("could not update user: user is empty")
	}
	if len(user) == 0 {
		return errors.New("user not found")
	}

	*u = user[0]
	return nil
}

func (u *User) AddListener(h UserHandler, ch chan string) {
	if u.listeners == nil {
		u.listeners = []UserHandler{}
	}
	if _, ok := u.session.listeners[u]; !ok {
		u.session.listeners[u] = []chan string{ch}
		u.listeners = append(u.listeners, h)
	} else {
		u.listeners = append(u.listeners, h)
	}
}

func (u *User) RemoveListener(l UserHandler, ch chan string) {
	for i := 0; i < len(u.listeners); i++ {
		if u.listeners[i] == l {
			
		}
	}
	/*if _, ok := s.listeners[u]; ok {
		for i := range s.listeners[e] {
			if s.listeners[e][i] == ch {
				s.listeners[e] = append(s.listeners[e][:i], s.listeners[e][i+1:]...)
				break
			}
		}
	}*/
}