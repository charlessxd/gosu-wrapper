package gosu

// Endpoint Constants
const (
	EndpointOsu               = "https://osu.ppy.sh/"
	EndpointAPI               = EndpointOsu + "api/"
	EndpointAPIKey            = "k"

	EndpointBeatmaps          = "get_beatmaps?"
	EndpointUser              = "get_user?"
	EndpointScores            = "get_scores?"
	EndpointUserBest          = "get_user_best?"
	EndpointUserRecent        = "get_user_recent?"
	EndpointMatch             = "get_match?"

	EndpointParamSince        = "since"
	EndpointParamBeatmapSetID = "s"
	EndpointParamBeatmapID    = "b"
	EndpointParamUserID       = "u"
	EndpointParamMode         = "m"
	EndpointParamType         = "type"
	EndpointParamConverted    = "a"
	EndpointParamHash         = "h"
	EndpointParamLimit        = "limit"
	EndpointParamMods   	  = "mods"
	EndpointParamMatchID      = "mp"
)
