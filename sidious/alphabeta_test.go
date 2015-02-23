package sidious

import (
	"github.com/danmane/deathstar/implgame"
	"math"
	"testing"
)

var s implgame.State = implgame.Standard

var surviveHeuristic Heuristic = func(g *implgame.State) int64 {
	return int64(g.NumPieces(implgame.White) - g.NumPieces(implgame.Black))
}
var suicideHeuristic Heuristic = func(g *implgame.State) int64 {
	return surviveHeuristic(g) * -1
}

func Test_AlphaBeta_SimpleSuicide(t *testing.T) {
	for depth := 1; depth < 3; depth++ {
		state := implgame.Standard
		oldstate := state
		lastNumPieces := state.NumPieces(implgame.White) + state.NumPieces(implgame.Black)
		for move := 0; move < 5; move++ {
			state = AlphaBetaWrap(&state, suicideHeuristic, depth)
			numPieces := state.NumPieces(implgame.White) + state.NumPieces(implgame.Black)
			if numPieces >= lastNumPieces {
				t.Errorf("suicide heuristic did not produce suicidal move!\ndepth:%v\n move:%v\n oldstate:%v\n state:%v\n",
					depth, move, oldstate, state)
			}
			lastNumPieces = numPieces
			oldstate = state
		}
	}
}

func Test_AlphaBeta_Survive(t *testing.T) {
	for depth := 1; depth < 3; depth++ {
		state := implgame.Standard
		oldstate := state
		lastNumPieces := state.NumPieces(implgame.White) + state.NumPieces(implgame.Black)
		for move := 0; move < 5; move++ {
			state = AlphaBetaWrap(&state, surviveHeuristic, depth)
			numPieces := state.NumPieces(implgame.White) + state.NumPieces(implgame.Black)
			if numPieces != lastNumPieces {
				t.Errorf("survival heuristic did produced suicidal move!\ndepth:%v\n move:%v\n oldstate:%v\n state:%v\n",
					depth, move, oldstate, state)
			}
			lastNumPieces = numPieces
			oldstate = state
		}
	}
}

func Test_AlphaBetaWrap_ReturnsAFuture(t *testing.T) {
	for i := 1; i < 2; i++ {
		f := AlphaBetaWrap(&s, DefaultHeuristic, i)
		if !s.ValidFuture(f) {
			t.Error("got an invalid future from AlphaBetaWrap", f, "when i=", i)
		}
	}
}

func Benchmark_AlphaBeta_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 1, math.MinInt64, math.MaxInt64, true, DefaultHeuristic)
	}
}

func Benchmark_AlphaBeta_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 2, math.MinInt64, math.MaxInt64, true, DefaultHeuristic)
	}
}

func Benchmark_AlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 3, math.MinInt64, math.MaxInt64, true, DefaultHeuristic)
	}
}

func Benchmark_TimedAlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		TimedAlphaBeta(&s, 3, true, 3*1000*1000*1000, DefaultHeuristic)
	}
}
