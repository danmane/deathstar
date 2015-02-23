package sidious

import (
	"fmt"
	"github.com/danmane/deathstar/implgame"
	"math"
	"time"
)

func AlphaBetaWrap(state *implgame.State, h Heuristic, depth int) (r implgame.State) {
	if depth < 1 {
		panic("invalid depth for alphabetawrap")
	}
	_, r = AlphaBeta(state, depth, math.MinInt64, math.MaxInt64, state.NextPlayer == implgame.White, h)
	return
}

func AlphaBeta(state *implgame.State, depth int, alpha, beta int64, maximizer bool, h Heuristic) (int64, implgame.State) {
	if depth == 0 || state.GameOver() {
		return h(state), *state
	}
	var bestVal int64
	var testVal int64
	var bestState implgame.State
	futures := SortedFutures(state)
	if maximizer {
		bestVal = math.MinInt64
		for _, f := range futures.states {
			testVal, _ = AlphaBeta(&f, depth-1, alpha, beta, false, h)
			if testVal > bestVal {
				bestVal = testVal
				bestState = f
				alpha = max64(alpha, bestVal)
				if beta <= alpha {
					break
				}
			}
		}
		return bestVal, bestState
	} else {
		bestVal = math.MaxInt64
		for _, f := range futures.states {
			testVal, _ = AlphaBeta(&f, depth-1, alpha, beta, true, h)
			if testVal < bestVal {
				bestVal = testVal
				bestState = f
				beta = min64(beta, bestVal)
				if beta <= alpha {
					break
				}
			}
		}
		return bestVal, bestState
	}
}

func TimedAlphaBeta(state *implgame.State, depth int, maximize bool, timeLimit int64, h Heuristic) implgame.State {
	var bestVal int64
	if maximize {
		bestVal = math.MinInt64
	} else {
		bestVal = math.MaxInt64
	}
	var bestState implgame.State
	futures := SortedFutures(state)

	timeChan := time.After(time.Duration(timeLimit - 1000*1000*5))

	for i := 0; i < futures.Len(); i++ {
		select {
		case <-timeChan:
			fmt.Println("breaking out of TAB at iteration %v of %v\n", i, futures.Len())
			return bestState
		default:
			future := futures.states[i]
			if maximize {
				testVal, _ := AlphaBeta(&future, depth-1, bestVal, math.MaxInt64, !maximize, h)
				if testVal > bestVal {
					bestVal = testVal
					bestState = future
				}
			} else {
				testVal, _ := AlphaBeta(&future, depth-1, math.MinInt64, bestVal, !maximize, h)
				if testVal < bestVal {
					bestVal = testVal
					bestState = future
				}

			}
		}
	}
	return bestState
}
