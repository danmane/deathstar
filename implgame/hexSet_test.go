package implgame

func makeHexSet(hexes ...Hex) HexSet {
	return slice2HexSet(hexes)
}
