package implgame

type Board struct {
	WhitePositions HexIndexArray `json:"whitePositions"`
	BlackPositions HexIndexArray `json:"blackPositions"`
	EdgeLength     int           `json:"edgeLength"`
}

func (b *Board) Pieces(p Player) HexIndexArray {
	if p == White {
		return b.WhitePositions
	} else if p == Black {
		return b.BlackPositions
	} else {
		var out HexIndexArray
		return out
	}
}

func (b *Board) free(x Hex) bool {
	return b.owner(x) == NullPlayer
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
	return x.Dist2Origin() < b.EdgeLength*2
}

var standardBoard Board = Board{
	EdgeLength:     5,
	WhitePositions: standardWhitePositions,
	BlackPositions: standardBlackPositions,
}
