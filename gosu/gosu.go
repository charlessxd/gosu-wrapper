// Package gosu provides a method of accessing osu-api in Go programs.
package gosu

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Session holds the API key and rate limiter.
type session struct {
	// Osu API Key
	key string

	// Rate limit
	limiter *RateLimit
}

// NewSession creates a Session using the user's APIKey.
func NewSession(APIKey string) (s session) {
	if APIKey == "" {
		return
	}

	s = session{
		key:     APIKey,
		limiter: NewRateLimit(),
	}

	return s
}

func (s *session) buildCall(endpoint string, v url.Values) string {
	return endpointAPI + endpoint + v.Encode()
}

func (s *session) parseJSON(url string, target interface{}) error {
	if !s.limiter.CanRequest {
		return errors.New("ratelimit exceeded")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	json.Unmarshal([]byte(buf.String()), &target)
	s.limiter.iterate()

	return err
}

func getMods(bitwiseDec int64) []string {
	binMods := ""
	for _, c := range strconv.FormatInt(bitwiseDec, 2) {
		binMods = string(c) + binMods
	}

	mods := []string{}

	if binMods == "0" {
		mods = append(mods, "No Mod")
	} else {
		if binMods[0] == '1' {
			mods = append(mods, "NF")
		}
		if len(binMods) >= 2 && binMods[1] == '1' {
			mods = append(mods, "EZ")
		}
		if len(binMods) >= 3 && binMods[2] == '1' {
			mods = append(mods, "TD")
		}
		if len(binMods) >= 4 && binMods[3] == '1' {
			mods = append(mods, "HD")
		}
		if len(binMods) >= 5 && binMods[4] == '1' {
			mods = append(mods, "HR")
		}
		if len(binMods) >= 6 && binMods[5] == '1' {
			if len(binMods) >= 15 && binMods[14] == '1' {
				mods = append(mods, "PF")
			} else {
				mods = append(mods, "SD")
			}
		}
		if len(binMods) >= 7 && binMods[6] == '1' {
			if len(binMods) >= 10 && binMods[9] == '1' {
				mods = append(mods, "NC")
			} else {
				mods = append(mods, "DT")
			}
		}
		if len(binMods) >= 8 && binMods[7] == '1' {
			mods = append(mods, "RL")
		}
		if len(binMods) >= 9 && binMods[8] == '1' {
			mods = append(mods, "HT")
		}
		if len(binMods) >= 11 && binMods[10] == '1' {
			mods = append(mods, "FL")
		}
		if len(binMods) >= 12 && binMods[11] == '1' {
			mods = append(mods, "Autoplay")
		}
		if len(binMods) >= 13 && binMods[12] == '1' {
			mods = append(mods, "SO")
		}
		if len(binMods) >= 14 && binMods[13] == '1' {
			mods = append(mods, "AP")
		}
		if len(binMods) >= 16 && binMods[15] == '1' {
			mods = append(mods, "4K")
		}
		if len(binMods) >= 17 && binMods[16] == '1' {
			mods = append(mods, "5K")
		}
		if len(binMods) >= 18 && binMods[17] == '1' {
			mods = append(mods, "6K")
		}
		if len(binMods) >= 19 && binMods[18] == '1' {
			mods = append(mods, "7K")
		}
		if len(binMods) >= 20 && binMods[19] == '1' {
			mods = append(mods, "8K")
		}
		if len(binMods) >= 21 && binMods[20] == '1' {
			mods = append(mods, "FI")
		}
		if len(binMods) >= 22 && binMods[21] == '1' {
			mods = append(mods, "RD")
		}
		if len(binMods) >= 23 && binMods[22] == '1' {
			mods = append(mods, "Cinema")
		}
		if len(binMods) >= 24 && binMods[23] == '1' {
			mods = append(mods, "TP")
		}
		if len(binMods) >= 25 && binMods[24] == '1' {
			mods = append(mods, "9K")
		}
		if len(binMods) >= 26 && binMods[25] == '1' {
			mods = append(mods, "KeyCoop")
		}
		if len(binMods) >= 27 && binMods[26] == '1' {
			mods = append(mods, "1K")
		}
		if len(binMods) >= 28 && binMods[27] == '1' {
			mods = append(mods, "3K")
		}
		if len(binMods) >= 29 && binMods[28] == '1' {
			mods = append(mods, "2K")
		}
		if len(binMods) >= 30 && binMods[29] == '1' {
			mods = append(mods, "ScoreV2")
		}
		if len(binMods) >= 31 && binMods[30] == '1' {
			mods = append(mods, "Mirror")
		}
	}
	return mods
}
