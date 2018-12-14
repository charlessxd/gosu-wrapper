package gosu

import (
	"strconv"
)

type UserHandler func(*Session, *User, User) bool

// PPGained is an nevent handler for when a user's PPRaw value changes.
func PPGained(s *Session, target *User, compare User) bool {
	if compare.PPRaw != target.PPRaw {
		s.Emit(*target, strconv.FormatFloat(compare.PPRaw-target.PPRaw, 'G', -1, 64))
		*target = compare
		return true
	}
	return false
}

// RankChange is an event handler for when a user's rank changes.
func RankChange(s *Session, target *User, compare User) bool {
	if compare.PPRank != target.PPRank {
		s.Emit(*target, strconv.FormatInt(compare.PPRank-target.PPRank, 64))
		*target = compare
		return true
	}
	return false
}