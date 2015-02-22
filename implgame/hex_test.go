package implgame

import "testing"

type conversionTest struct {
	q int
	r int
	i int
}

var conversionTests = []conversionTest{
	{q: 0, r: -4, i: 0},
	{q: 4, r: -4, i: 4},
	{q: -1, r: -3, i: 5},
	{q: 4, r: 0, i: 34},
	{q: 0, r: 4, i: 60},
}

func Test_toIdx(t *testing.T) {
	for _, test := range conversionTests {
		h := Hex{Q: test.q, R: test.r}
		if h.idx() != test.i {
			t.Error("toIdx: q:", h.Q, ", r:", h.R, "expected ", test.i, "got", h.idx())
		}
	}
}

func Test_fromIdx(t *testing.T) {
	for _, test := range conversionTests {
		hex := fromIdx(test.i)
		if hex.idx() != test.i {
			t.Error("idx", test.i,
				"expected qr", test.q, test.r,
				"got qr", hex.Q, hex.R)
		}
	}
}

func Benchmark_Hex_toIdx(b *testing.B) {
	var allHexes [61]Hex
	for i := 0; i < 61; i++ {
		allHexes[i] = *fromIdx(i)
	}
	for n := 0; n <= b.N; n++ {
		for i := 0; i < 61; i++ {
			allHexes[i].idx()
		}
	}
}
func Benchmark_Hex_fromIdx(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		for i := 0; i < 61; i++ {
			fromIdx(i)
		}
	}
}
