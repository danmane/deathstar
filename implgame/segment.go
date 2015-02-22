package implgame

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

func (g *State) Segments(p Player) []Segment {
	pieces := g.Board.Pieces(p)
	result := make([]Segment, 0, 3*len(pieces))
	for _, pos := range pieces.ToSlice() {
		s := Segment{
			base:        pos,
			Length:      1,
			player:      p,
			orientation: NullDirection,
		}
		result = append(result, s)
		for d := TopRight; d <= BotRight; d++ {
			next := pos.adjacent(d)
			length := 2
			for length <= g.MarblesPerMove && pieces.has(next) {
				s = Segment{
					base:        pos,
					orientation: d,
					Length:      length,
					player:      p,
				}
				next = next.adjacent(d)
				length++
				result = append(result, s)
			}
		}
	}
	return result
}
