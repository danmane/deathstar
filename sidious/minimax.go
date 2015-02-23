package sidious

import (
	"github.com/danmane/deathstar/implgame"
	"math"
)

func Minimax(state *implgame.State, depth int, maximizer bool, weights HeuristicWeights) (int64, implgame.State) {
	if depth == 0 || state.GameOver() {
		return CalcHeuristic(state, weights), *state
	}
	var bestVal int64
	var testVal int64
	var bestState implgame.State
	futures := state.Futures()
	if maximizer {
		bestVal = math.MinInt64
		for _, f := range futures {
			testVal, _ = Minimax(&f, depth-1, false, weights)
			if testVal > bestVal {
				bestVal = testVal
				bestState = f
			}
		}
		return bestVal, bestState
	} else {
		bestVal = math.MaxInt64
		for _, f := range futures {
			testVal, _ = Minimax(&f, depth-1, true, weights)
			if testVal < bestVal {
				bestVal = testVal
				bestState = f
			}
		}
		return bestVal, bestState
	}
}
