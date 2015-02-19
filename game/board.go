package game

type Board struct {
	WhitePositions HexSet `json:"whitePositions"`
	BlackPositions HexSet `json:"blackPositions"`
	EdgeLength     int    `json:"edgeLength"`
}

func (b *Board) pieces(p Player) HexSet {
	if p == White {
		return b.WhitePositions
	} else if p == Black {
		return b.BlackPositions
	} else {
		return make(HexSet)
	}
}

func (b *Board) free(x Hex) bool {
	return b.owner(x) == NullPlayer
}

func (b1 *Board) eq(b2 Board) bool {
	return b1.WhitePositions.eq(b2.WhitePositions) &&
		b1.BlackPositions.eq(b2.BlackPositions) &&
		b1.EdgeLength == b2.EdgeLength
}

func (b *Board) owner(x Hex) Player {
	if b.WhitePositions.has(x) {
		return White
	} else if b.BlackPositions.has(x) {
		return Black
	} else {
		return NullPlayer
	}
}

func (b *Board) onBoard(x Hex) bool {
	return x.dist2Origin() < b.EdgeLength*2
}

var standardBoard Board = Board{
	EdgeLength: 5,
	WhitePositions: slice2HexSet(
		[]Hex{
			{-4, 3}, {-4, 4}, {-3, 3}, {-3, 4}, {-2, 2},
			{-2, 3}, {-2, 4}, {-1, 2}, {-1, 3}, {-1, 4},
			{0, 2}, {0, 3}, {0, 4}, {1, 3},
		},
	),
	BlackPositions: slice2HexSet(
		[]Hex{
			{-1, -3}, {0, -4}, {0, -3}, {0, -2}, {1, -4},
			{1, -3}, {1, -2}, {2, -4}, {2, -3}, {2, -2},
			{3, -4}, {3, -3}, {4, -4}, {4, -3},
		},
	),
}
