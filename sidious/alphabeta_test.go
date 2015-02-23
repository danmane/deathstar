package sidious

import (
	"github.com/danmane/deathstar/implgame"
	"math"
	"testing"
)

var s implgame.State = implgame.Standard

func Test_AlphaBetaWrap_ReturnsAFuture(t *testing.T) {
	for i := 1; i < 2; i++ {
		f := AlphaBetaWrap(&s, DefaultWeights, i)
		if !s.ValidFuture(f) {
			t.Error("got an invalid future from AlphaBetaWrap", f, "when i=", i)
		}
	}
}

func Benchmark_AlphaBeta_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 1, math.MinInt64, math.MaxInt64, true, DefaultWeights)
	}
}

func Benchmark_AlphaBeta_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 2, math.MinInt64, math.MaxInt64, true, DefaultWeights)
	}
}

func Benchmark_AlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 3, math.MinInt64, math.MaxInt64, true, DefaultWeights)
	}
}

func Benchmark_TimedAlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		TimedAlphaBeta(&s, 3, true, 3*1000*1000*1000, DefaultWeights)
	}
}
