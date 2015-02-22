package implgame

type State struct {
	Board          Board  `json:"board"`
	NextPlayer     Player `json:"nextPlayer"`
	MovesRemaining int    `json:"movesRemaining"`
	MarblesPerMove int    `json:"marblesPerMove"`
	LossThreshold  int    `json:"lossThreshold"`
}

var Standard State = State{
	Board:          standardBoard,
	NextPlayer:     White,
	MovesRemaining: 1000,
	MarblesPerMove: 3,
	LossThreshold:  8,
}

func (g *State) Moves() []Move {
	result := make([]Move, 0)
	for _, s := range g.Segments(g.NextPlayer) {
		for _, d := range Directions {
			m := Move{Segment: s, direction: d}
			if m.isValid(g) {
				result = append(result, m)
			}
		}
	}
	return result
}

// like all functions in this implementation, this returns a copy
// if given an invalid move, behavior is undefined
func (g *State) Update(m *Move) State {
	ownPieces := m.Segment.segPieces()
	var enemyPieces []Hex
	if m.inline() {
		enemyPieces = m.inlineMoved(g.Board)
	} else {
		enemyPieces = make([]Hex, 0)
	}
	var whiteMoved, blackMoved []Hex
	if g.NextPlayer == White {
		whiteMoved = ownPieces
		blackMoved = enemyPieces
	} else {
		whiteMoved = enemyPieces
		blackMoved = ownPieces
	}

	copyAndMove := func(original HexIndexArray, hexesToMove []Hex) HexIndexArray {
		var out HexIndexArray
		for i := 0; i < 61; i++ {
			out[i] = original[i]
		}
		for _, h := range hexesToMove {
			out[h.idx()] = false
		}
		for _, h := range hexesToMove {
			adj := h.adjacent(m.direction)
			if g.Board.onBoard(adj) {
				out[adj.idx()] = true
			}
		}
		return out
	}

	newWhite := copyAndMove(g.Board.WhitePositions, whiteMoved)
	newBlack := copyAndMove(g.Board.BlackPositions, blackMoved)
	newBoard := Board{
		WhitePositions: newWhite,
		BlackPositions: newBlack,
		EdgeLength:     g.Board.EdgeLength,
	}
	newGame := State{
		Board:          newBoard,
		NextPlayer:     g.NextPlayer.Next(),
		MovesRemaining: g.MovesRemaining - 1,
		MarblesPerMove: g.MarblesPerMove,
		LossThreshold:  g.LossThreshold,
	}

	return newGame
}

func (g *State) Futures() []State {
	moves := g.Moves()
	result := make([]State, len(moves))
	for i := 0; i < len(moves); i++ {
		result[i] = g.Update(&moves[i])
	}
	return result
}

func (g1 *State) ValidFuture(g2 State) bool {
	for _, future := range g1.Futures() {
		if future == g2 {
			return true
		}
	}
	return false
}

func (g *State) GameOver() bool {
	return g.Outcome() != NullOutcome
}

func (g *State) NumPieces(p Player) int {
	return len(g.Board.Pieces(p))
}

func (g *State) Outcome() Outcome {
	w := len(g.Board.WhitePositions)
	b := len(g.Board.BlackPositions)
	if g.MovesRemaining <= 0 || w <= g.LossThreshold || b <= g.LossThreshold {
		if w < b {
			return BlackWins
		} else if b < w {
			return WhiteWins
		} else {
			return Tie
		}
	} else {
		return NullOutcome
	}
}
