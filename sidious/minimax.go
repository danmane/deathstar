package sidious

import (
	"github.com/danmane/sidious/implgame"
	"math"
)

func Minimax(state *implgame.State, depth int, maximizer bool) (float64, *implgame.State) {
	if depth == 0 || state.GameOver() {
		return myHeuristic(state), state
	}
	var bestVal float64
	var testVal float64
	var bestState *implgame.State
	futures := state.Futures()
	if maximizer {
		bestVal = math.Inf(-1)
		for _, f := range futures {
			testVal, _ = Minimax(&f, depth-1, false)
			if testVal > bestVal {
				bestVal = testVal
				bestState = &f
			}
		}
		return bestVal, bestState
	} else {
		bestVal = math.Inf(1)
		for _, f := range futures {
			testVal, _ = Minimax(&f, depth-1, true)
			if testVal < bestVal {
				bestVal = testVal
				bestState = &f
			}
		}
		return bestVal, bestState
	}
}

func myHeuristic(state *implgame.State) float64 {
	if state.GameOver() {
		switch state.Outcome() {
		case implgame.WhiteWins:
			return math.Inf(1)
		case implgame.BlackWins:
			return math.Inf(-1)
		default:
			return 0.0
		}
	}
	weights := []float64{10.0, 1.0, 1.0, 1.0, 1.0}
	var val float64 = 0
	for i, _ := range Heuristics {
		val += float64(Heuristics[i](state, implgame.White)) * weights[i]
	}
	return val
}
