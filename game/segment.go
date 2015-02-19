package game

type Segment struct {
	base        Hex
	Length      int
	player      Player
	orientation Direction
}

func (s *Segment) segPieces() []Hex {
	result := make([]Hex, s.Length)
	p := s.base
	for i := 0; i < s.Length; i++ {
		result[i] = p
		p = p.adjacent(s.orientation)
	}
	return result
}
