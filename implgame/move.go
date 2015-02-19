package implgame

type Move struct {
	Segment   Segment
	direction Direction
}

func (m *Move) isValid(g *State) bool {
	if m.inline() {
		return m.inlineMoved(g.Board) != nil
	} else {
		for _, h := range m.Segment.segPieces() {
			dest := h.adjacent(m.direction)
			if !g.Board.free(dest) {
				return false
			}
		}
		return true
	}
}

func (m *Move) inline() bool {
	return m.direction.colinear(m.Segment.orientation)
}

func (m *Move) inlineMoved(b Board) []Hex {
	movedEnemyPieces := make([]Hex, 0)
	pieces := m.Segment.segPieces()
	var attacked Hex
	if m.Segment.orientation == m.direction {
		attacked = pieces[len(pieces)-1]
	} else {
		attacked = pieces[0]
	}
	for i := 0; i < m.Segment.Length; i++ {
		attacked = attacked.adjacent(m.direction)
		controller := b.owner(attacked)
		if controller == NullPlayer {
			return movedEnemyPieces
		} else if controller == m.Segment.player {
			return nil
		}
		movedEnemyPieces = append(movedEnemyPieces, attacked)
	}
	return nil
}
