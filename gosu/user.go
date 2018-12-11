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
	Count300 string `json:"count300"`

	// The number of 100s the user has gotten for all ranked, approved, and loved beatmaps played.
	Count100 string `json:"count100"`

	// The number of 50s the user has gotten for all ranked, approved, and loved beatmaps played.
	Count50 string `json:"count50"`

	// The amount of plays the user has on ranked, approved, and loved beatmaps.
	PlayCount string `json:"playcount"`

	// The total of the best individual score for every ranked, approved, and loved beatmap the user has played.
	RankedScore string `json:"ranked_score"`

	// The total score of every beatmap the user has played.
	TotalScore string `json:"total_score"`

	// The global ranking of the user in terms of PP.
	// 1 having the highest amount of PP.
	PPRank string `json:"pp_rank"`

	// The level of the user.
	Level string `json:"level"`

	// The amount of PP the user has as a float value.
	PPRaw string `json:"raw"`

	// The accuracy of the user as a percentage.
	Accuracy string `json:"accuracy"`

	// The total number of non-hidden SS ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankSS string `json:"count_rank_ss"`

	// The total number of hidden SS ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankSSH string `json:"count_rank_ssh"`

	// The total number of non-hidden S ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankS string `json:"count_rank_s"`

	// The total number of hidden S ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankSH string `json:"count_rank_sh"`

	// The total number of A ranks the user has achieved on ranked, approved, and loved beatmaps.
	CountRankA string `json:"count_rank_a"`

	// The country the user created their account from.
	// Uses the ISO3166-1 alpha-2 country code naming. See this for more information: http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2/)
	Country string `json:"country"`

	// The total number of seconds the user has played.
	TotalSecondsPlayed string `json:"total_seconds_played"`

	// The PP ranking of the user within their country.
	PPCountryRank string `json:"pp_country_rank"`

	// The events of the user.
	Events []UserEvent `json:"events"`
}

// UserEvent stores data for events related to an individual osu user.
type UserEvent []struct {
	// The HTML for the event.
	DisplayHTML string `json:"display_html"`

	// The ID of the beatmap in the event.
	BeatmapID string `json:"beatmap_id"`

	// The ID of beatmap set in the event.
	BeatmapSetID string `json:"beatmapset_id"`

	// The date the event occurred, in UTC.
	Date string `json:"date"`

	// How epic the event is, between 1 and 32.
	EpicFactor string `json:"epicfactor"`
}

// FetchUser returns metadata about a user.
func (s *Session) FetchUser(call UserCall) (User, error) {
	user := new([]User)
	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)

	switch {
	case call.UserID != "":
		v.Add(EndpointUserUserID, call.UserID)
	default:
		return User{}, errors.New("no identifying parameter given (UserID)")
	}

	if call.Mode != "" {
		v.Add(EndpointUserMode, call.Mode)
	}
	if call.Type != "" {
		v.Add(EndpointUserType, call.Type)
	}

	s.ParseJSON(s.BuildCall(EndpointUser, v), user)

	if len(*user) == 0 {
		return User{}, errors.New("user not found")
	}

	return (*user)[0], nil
}
