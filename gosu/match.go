package gosu

import (
	"errors"
	"net/url"
)

// MatchCall is used to build an API call to retrieve metadata on a match.
type MatchCall struct {
	// ID of target match.
	MatchID string
}

// Match stores data about a specific multi-player match.
type Match struct {
	// Details of the match.
	Details MatchDetails `json:"match"`

	// Contains all of the games played in the match.
	Games []MatchGame `json:"games"`

	// API Call URL.
	apiURL string

	// Session fetched from
	session *Session
}

// MatchDetails stores data for the details of a specific multi-player match.
type MatchDetails struct {
	// The ID of the match.
	MatchID string `json:"match_id"`

	// The name of the match.
	Name string `json:"name"`

	// The time when the match was created, in UTC.
	StartTime string `json:"start_time"`

	// The time when the match ended, in UTC.
	EndTime string `json:"end_time"`
}

// MatchGame stores data for all of the games played in a specific multi-player match.
type MatchGame struct {
	// The ID of the game.
	GameID string `json:"game_id"`

	// The time when the game started, in UTC.
	StartTime string `json:"start_time"`

	// The time when the game ended, in UTC.
	EndTime string `json:"end_time"`

	// The ID of the beatmap played in the game.
	BeatmapID string `json:"beatmap_id"`

	// The game mode played in the game.
	// standard = 0, taiko = 1, ctb = 2, mania = 3
	PlayMode string `json:"play_mode"`

	// The type of match.
	MatchType string `json:"match_type"`

	// The type of scoring used for the game.
	// score = 0, accuracy = 1, combo = 2, score v2 = 3
	ScoringType string `json:"scoring_type"`

	// The bitwise flag representation of the mods used.
	Mods string `json:"mods"`

	// The scores of all users who participated in the game.
	Scores []MatchScore `json:"scores"`
}

// MatchScore stores data for each individual user who participated in a game.
type MatchScore struct {
	// Zero-based index of the user's slot.
	Slot string `json:"slot"`

	// What team the user was on.
	//  0 = mode doesn't support teams, 1 = blue, 2 = red
	Team string `json:"team"`

	// The ID of the user.
	UserID string `json:"user_id"`

	// The score achieved.
	Score string `json:"score"`

	// The highest combo the user reached.
	MaxCombo string `json:"maxcombo"`

	// Not used.
	Rank string `json:"rank"`

	// The number of 50s.
	Count50 string `json:"count50"`

	// The number of 100s.
	Count100 string `json:"count100"`

	// The number of 300s.
	Count300 string `json:"count300"`

	// The number of misses.
	CountMiss string `json:"coutmiss"`

	// The number of geki.
	CountGeki string `json:"countgeki"`

	// The number of katu.
	CountKatu string `json:"countkatu"`

	// Whether the user achieved the maximum combo of the beatmap.
	// 1 = max combo achieved, 0 = max combo not achieved.
	Perfect string `json:"perfect"`

	// Whether or not the user passed the beatmap.
	// 1 = was not failed when the beatmap ended, 0 = was failed when the beatmap ended.
	Pass string `json:"pass"`
}

// FetchMatch returns metadata about a match.
func (s *Session) FetchMatch(call MatchCall) (Match, error) {
	match := *new([]Match)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.MatchID != "":
		v.Add(endpointParamMatchID, call.MatchID)
	default:
		return Match{}, errors.New("no identifying param given (MatchID)")
	}

	err := s.parseJSON(s.buildCall(endpointMatch, v), match)

	if err != nil {
		return Match{}, err
	}
	if len(match) == 0 {
		return Match{}, errors.New("match not found")
	}

	match[0].apiURL = s.buildCall(endpointMatch, v)
	match[0].session = s

	return match[0], nil
}

// Updates a Match.
func (m *Match) Update() error {
	match := *new([]Match)

	err := m.session.parseJSON(m.apiURL, match)

	if err != nil {
		return err
	}
	if m.apiURL == "" {
		return errors.New("could not update user: user is empty")
	}
	if len(match) == 0 {
		return errors.New("user not found")
	}

	*m = match[0]
	return nil
}
