package gosu

// Endpoint Constants
const (
	EndpointOsu    = "https://osu.ppy.sh/"
	EndpointAPI    = EndpointOsu + "api/"
	EndpointAPIKey = "k"

	EndpointBeatmaps             = "get_beatmaps?"
	EndpointBeatmapsSince        = "since"
	EndpointBeatmapsBeatmapSetID = "s"
	EndpointBeatmapsBeatmapID    = "b"
	EndpointBeatmapsUserID       = "u"
	EndpointBeatmapsMode         = "m"
	EndpointBeatmapsType         = "type"
	EndpointBeatmapsConverted    = "a"
	EndpointBeatmapsHash         = "h"
	EndpointBeatmapsLimit        = "limit"

	EndpointUser       = EndpointAPI + "get_user?"
	EndpointUserUserID = "u"
	EndpointUserMode   = "m"
	EndpointUserType   = "type"

	EndpointScores          = EndpointAPI + "get_scores?"
	EndpointScoresBeatmapID = "b"
	EndpointScoresUserID    = "u"
	EndpointScoresMode      = "m"
	EndpointScoresMods      = "mods"
	EndpointScoresType      = "type"
	EndpointScoresLimit     = "limit"

	EndpointUserBest       = EndpointAPI + "get_user_best?"
	EndpointUserBestUserID = "u"
	EndpointUserBestMode   = "m"
	EndpointUserBestType   = "type"
	EndpointUserBestLimit  = "limit"

	EndpointUserRecent       = EndpointAPI + "get_user_recent?"
	EndpointUserRecentUserID = "u"
	EndpointUserRecentMode   = "m"
	EndpointUserRecentType   = "type"
	EndpointUserRecentLimit  = "limit"

	EndpointMatch   = "get_match?"
	EndpointMatchID = "mp"
)
