package game

import "github.com/eltrufas/tetriscore"

type State struct {
	T                *tetriscore.Tetris
	OffsetX, OffsetY int
	BlockW           int
}

func CreateGame() *State {
	var s State
	s.OffsetX = 100
	s.OffsetY = 470
	s.BlockW = 16
	s.T = tetriscore.CreateTetris()
	return &s
}
