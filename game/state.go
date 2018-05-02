package game

import "github.com/eltrufas/tetriscore"
import "github.com/eltrufas/rltetris"

type State struct {
	T                *tetriscore.Tetris
	OffsetX, OffsetY int
	BlockW           int
	Weights          map[uint32][]float64
	Action           uint32
	lastAction       tetriscore.InputState
	player           *rltetris.RemotePlayer
}

func CreateGame() *State {
	var s State
	s.OffsetX = 100
	s.OffsetY = 470
	s.BlockW = 16
	s.T = tetriscore.CreateTetris()
	s.player = rltetris.CreateRemotePlayer("localhost:5050")

	return &s
}
