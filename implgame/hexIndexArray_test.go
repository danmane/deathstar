package implgame

import "testing"

func makeHexIndexArray(hexes ...Hex) HexIndexArray {
	return slice2HexIndexArray(hexes)
}

func Benchmark_HexIndexArray_lookup(b *testing.B) {
	allHexes := make([]Hex, 61)
	for i := 0; i < 61; i++ {
		allHexes[i] = *fromIdx(i)
	}
	fullSet := slice2HexIndexArray(allHexes)
	for n := 0; n <= b.N; n++ {
		for i := 0; i < 61; i++ {

			fullSet.has(allHexes[i])
		}
	}
}

func Test_conversion(t *testing.T) {
	converted := slice2HexIndexArray(standardWhitePositions.toSlice())
	if standardWhitePositions != converted {
		t.Log("something went wrong in HexIndexArray conversion")
		for i := 0; i < 61; i++ {
			if standardWhitePositions[i] != converted[i] {
				t.Logf("original[%v] = %v", i, standardWhitePositions[i])
			}
		}
		t.Error("conversion from", standardWhitePositions, "to", converted, "failed")
	}
}
