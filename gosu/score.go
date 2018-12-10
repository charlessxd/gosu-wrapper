package gosu

import (
	"errors"
	"net/url"
)

// ScoreCall is used to build an API call to retrieve metadata of scores set on a beatmap.
type ScoreCall struct {
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

// Scores stores the data for the top 100 scores of a specific beatmap.
type Scores []struct {
	// The ID of the score.
	ScoreID string `json:"score_id"`

	// The score achieved.
	Score string `json:"score"`

	// The user name of the user who submitted the score.
	Username string `json:"username"`

	// The number of 300s.
	Count300 string `json:"count300"`

	// The number of 100s.
	Count100 string `json:"count100"`

	// The number of 50s.
	Count50 string `json:"count50"`

	// The number of misses.
	CountMiss string `json:"countmiss"`

	// The highest combo the user reached.
	MaxCombo string `json:"maxcombo"`

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

	// Whether osu official servers store the replay.
	// 1 = is stored, 0 = is not stored.
	ReplayAvailable string `json:"replay_available"`
}

// FetchScores returns metadata about scores set on a beatmap.
func (s *Session) FetchScores(call *ScoreCall) (Scores, error) {
	scores := new(Scores)
	v := url.Values{}
	v.Add(EndpointAPIKey, s.Key)

	switch {
	case call.BeatmapID != "":
		v.Add(EndpointScoresBeatmapID, call.BeatmapID)
	default:
		return Scores{}, errors.New("no identifying parameter given (BeatmapID)")
	}

	if call.UserID != "" {
		v.Add(EndpointScoresUserID, call.UserID)
	}
	if call.Mode != "" {
		v.Add(EndpointScoresMode, call.Mode)
	}
	if call.Mods != "" {
		v.Add(EndpointScoresMods, call.Mods)
	}
	if call.Type != "" {
		v.Add(EndpointScoresType, call.Type)
	}
	if call.Limit != "" {
		v.Add(EndpointScoresLimit, call.Limit)
	}

	s.ParseJSON(s.BuildCall(EndpointScores, v), scores)

	if len(*scores) == 0 {
		return *scores, errors.New("no scores found")
	}

	return *scores, nil
}
