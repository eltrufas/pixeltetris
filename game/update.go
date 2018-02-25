package game

import "github.com/eltrufas/tetriscore"
import "fmt"

func (s *State) Update(ia []bool) bool {
	var is tetriscore.InputState
	for i, input := range ia {
		if input {
			is |= 1 << uint32(i)
		}
	}

	s.T.Update(is)
	return true
}
