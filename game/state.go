package game

import "github.com/eltrufas/tetriscore"
import "github.com/eltrufas/rltetris"

type State struct {
	T                *tetriscore.Tetris
	OffsetX, OffsetY int
	BlockW           int
  Weights          map[uint32][]float64
  Game             rltetris.Tetrisrl
  Action           uint32
  ShouldPress      bool
  Q                [8]float64
}

func CreateGame() *State {
	var s State
	s.OffsetX = 100
	s.OffsetY = 470
	s.BlockW = 16
  s.Game = rltetris.CreateTetris()
  s.T = s.Game.Tetris
  s.Weights = make(map[uint32][]float64)
  feats := len(s.Game.GetState())
  for _, a := range s.Game.LegalAction() {
    s.Weights[a] = make([]float64, feats)
  }

  //rltetris.Sarsa(s.Weights, 1000, 0.3, 1)
	return &s
}
