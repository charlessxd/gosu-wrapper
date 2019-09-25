package gosu

import (
	"errors"
	"net/url"
	"strconv"
)

// ScoresCall is used to build an API call to retrieve metadata of scores set on a beatmap.
type ScoresCall struct {
	// ID of the beatmap
	BeatmapID string

	// ID or Username of the target user.
	UserID string

	// Specific game-mode.
	// 0 = standard, 1 = taiko, 2 = ctb, 3 = mania
	Mode string

	// Specific mod combination
	Mods string

	// Whether UserID is an ID or a Username.
	// "id" if ID
	// "string" if username
	Type string

	// Maximum amount of results
	// Range between 1 - 100 (defaults to 10).
	Limit string
}

// Score stores the data for the top 100 scores of a specific beatmap.
type Score struct {
	// The ID of the score.
	ScoreID string `json:"score_id"`

	// The score achieved.
	Score int64 `json:"score,string"`

	// The user name of the user who submitted the score.
	Username string `json:"username"`

	// The number of 300s.
	Count300 int `json:"count300,string"`

	// The number of 100s.
	Count100 int `json:"count100,string"`

	// The number of 50s.
	Count50 int `json:"count50,string"`

	// The number of misses.
	CountMiss int `json:"countmiss,string"`

	// The highest combo the user reached.
	MaxCombo int `json:"maxcombo,string"`

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

	// Whether osu official servers store the replay.
	// 1 = is stored, 0 = is not stored.
	ReplayAvailable string `json:"replay_available"`

	Accuracy float64
}

// Scores stores scores.
type Scores struct {
	Scores []Score

	// API Call URL.
	apiURL string

	// Session fetched from
	session *session

	apiCall ScoresCall
}

// FetchScores returns metadata about scores set on a beatmap.
func (s *session) FetchScores(call ScoresCall) (Scores, error) {
	if i, e := strconv.ParseInt(call.Limit, 10, 64); i > 100 {
		return Scores{}, errors.New("limit parameter exceeds maximum limit (Max limit: 100)")
	} else if e != nil {
		return Scores{}, e
	}

	scores := *new([]Score)
	v := url.Values{}
	v.Add(endpointAPIKey, s.key)

	switch {
	case call.BeatmapID != "":
		v.Add(endpointParamBeatmapID, call.BeatmapID)
	default:
		return Scores{}, errors.New("no identifying parameter given (BeatmapID)")
	}

	if call.UserID != "" {
		v.Add(endpointParamUserID, call.UserID)
	}
	if call.Mode != "" {
		v.Add(endpointParamMode, call.Mode)
	}
	if call.Mods != "" {
		v.Add(endpointParamMods, call.Mods)
	}
	if call.Type != "" {
		v.Add(endpointParamType, call.Type)
	}
	if call.Limit != "" {
		v.Add(endpointParamLimit, call.Limit)
	}

	err := s.parseJSON(s.buildCall(endpointScores, v), &scores)

	ss := *new(Scores)
	ss.Scores = scores

	if err != nil {
		return ss, err
	}
	if len(scores) == 0 {
		return ss, errors.New("no scores found")
	}

	for i, score := range ss.Scores {
		ss.Scores[i].EnabledMods = getMods(score.ModsInt)
		ss.Scores[i].Accuracy = getAcc(score.CountMiss, score.Count50, score.Count100, score.Count300)
	}

	ss.apiURL = s.buildCall(endpointScores, v)
	ss.session = s

	return ss, nil
}

// Update updates the Scores on a Beatmap
func (ss *Scores) Update() error {
	temp, err := ss.session.FetchScores(ss.apiCall)
	*ss = temp

	if err != nil {
		return err
	}
	return nil
}
