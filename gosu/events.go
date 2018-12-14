package gosu

import (
	"strconv"
)

type UserHandler func(*Session, *User, User) bool

func PPGained(s *Session, target *User, compare User) bool {
	if compare.PPRaw > target.PPRaw {
		s.Emit(*target, strconv.FormatFloat(compare.PPRaw-target.PPRaw, 'G', -1, 64))
		*target = compare
		return true
	}
	return false
}