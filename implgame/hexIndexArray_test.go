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
	converted := slice2HexIndexArray(standardWhitePositions.ToSlice())
	if standardWhitePositions != converted {
		t.Log("something went wrong in HexIndexArray conversion")
		for i := 0; i < 61; i++ {
			if standardWhitePositions.arr[i] != converted.arr[i] {
				t.Logf("original[%v] = %v", i, standardWhitePositions.arr[i])
			}
		}
		t.Error("conversion from", standardWhitePositions, "to", converted, "failed")
	}
}

func hia2int(arr [61]bool) uint64 {
	var out uint64 = 0
	var i uint64
	for i = 0; i < 61; i++ {
		if arr[i] {
			out |= (1 << i)
		}
	}
	return out
}

func Benchmark_HIA_copy(b *testing.B) {
	var CopiedHIA [61]bool
	a := standardWhitePositions.arr
	for n := 0; n <= b.N; n++ {
		CopiedHIA = a
	}
	a = CopiedHIA
}

func Benchmark_int_copy(b *testing.B) {
	var CopiedInt uint64
	a := hia2int(standardWhitePositions.arr)
	for n := 0; n <= b.N; n++ {
		CopiedInt = a
	}
	a = CopiedInt
}

func Benchmark_HIA_ToSlice(b *testing.B) {
	a := standardWhitePositions
	for n := 0; n <= b.N; n++ {
		a.ToSlice()
	}
}

func Benchmark_int_ToSlice(b *testing.B) {
	a := hia2int(standardWhitePositions.arr)
	for n := 0; n <= b.N; n++ {
		out := make([]Hex, standardWhitePositions.length)
		x := 0
		var i uint64 = 0
		for i = 0; i < 61; i++ {
			if (a>>i)&1 == 1 {
				out[x] = *fromIdx(int(i))
				x++
			}
		}
	}
}
