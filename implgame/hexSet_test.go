package implgame

import "testing"

func makeHexSet(hexes ...Hex) HexSet {
	return slice2HexSet(hexes)
}

func Benchmark_hexSet_lookup(b *testing.B) {
	allHexes := make([]Hex, 61)
	for i := 0; i < 61; i++ {
		allHexes[i] = *fromIdx(i)
	}
	fullSet := slice2HexSet(allHexes)
	for n := 0; n <= b.N; n++ {
		for i := 0; i < 61; i++ {

			fullSet.has(allHexes[i])
		}
	}
}
