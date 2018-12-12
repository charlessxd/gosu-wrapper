package gosu

// endpoint Constants
const (
	endpointOsu    = "https://osu.ppy.sh/"
	endpointAPI    = endpointOsu + "api/"
	endpointAPIKey = "k"

	endpointBeatmaps   = "get_beatmaps?"
	endpointUser       = "get_user?"
	endpointScores     = "get_scores?"
	endpointUserBest   = "get_user_best?"
	endpointUserRecent = "get_user_recent?"
	endpointMatch      = "get_match?"

	endpointParamSince        = "since"
	endpointParamBeatmapSetID = "s"
	endpointParamBeatmapID    = "b"
	endpointParamUserID       = "u"
	endpointParamMode         = "m"
	endpointParamType         = "type"
	endpointParamConverted    = "a"
	endpointParamHash         = "h"
	endpointParamLimit        = "limit"
	endpointParamMods         = "mods"
	endpointParamMatchID      = "mp"
)
