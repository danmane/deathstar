package main

import (
	"fmt"
	"github.com/danmane/deathstar/implgame"
	"math"
	"time"
)

func Minimax(state *implgame.State, depth int, maximizer bool) (float64, implgame.State) {
	if depth == 0 || state.GameOver() {
		return myHeuristic(state, implgame.White), *state
	}
	var bestVal float64
	var testVal float64
	var bestState implgame.State
	futures := state.Futures()
	if maximizer {
		bestVal = math.Inf(-1)
		for _, f := range futures {
			testVal, _ = Minimax(&f, depth-1, false)
			if testVal > bestVal {
				bestVal = testVal
				bestState = f
			}
		}
		return bestVal, bestState
	} else {
		bestVal = math.Inf(1)
		for _, f := range futures {
			testVal, _ = Minimax(&f, depth-1, true)
			if testVal < bestVal {
				bestVal = testVal
				bestState = f
			}
		}
		return bestVal, bestState
	}
}

func TimedAlphaBeta(state *implgame.State, depth int, maximize bool, timeLimit int64) implgame.State {
	var bestVal float64
	if maximize {
		bestVal = math.Inf(-1)
	} else {
		bestVal = math.Inf(1)
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
				testVal, _ := AlphaBeta(&future, depth-1, bestVal, math.Inf(1), !maximize)
				if testVal > bestVal {
					bestVal = testVal
					bestState = future
				}
			} else {
				testVal, _ := AlphaBeta(&future, depth-1, math.Inf(-1), bestVal, !maximize)
				if testVal < bestVal {
					bestVal = testVal
					bestState = future
				}

			}
		}
	}
	return bestState
}

func AlphaBeta(state *implgame.State, depth int, alpha, beta float64, maximizer bool) (float64, implgame.State) {
	if depth == 0 || state.GameOver() {
		return myHeuristic(state, implgame.White), *state
	}
	var bestVal float64
	var testVal float64
	var bestState implgame.State
	futures := SortedFutures(state)
	if maximizer {
		bestVal = math.Inf(-1)
		for _, f := range futures.states {
			testVal, _ = AlphaBeta(&f, depth-1, alpha, beta, false)
			if testVal > bestVal {
				bestVal = testVal
				bestState = f
				alpha = math.Max(alpha, bestVal)
				if beta <= alpha {
					break
				}
			}
		}
		return bestVal, bestState
	} else {
		bestVal = math.Inf(1)
		for _, f := range futures.states {
			testVal, _ = AlphaBeta(&f, depth-1, alpha, beta, true)
			if testVal < bestVal {
				bestVal = testVal
				bestState = f
				beta = math.Min(beta, bestVal)
				if beta <= alpha {
					break
				}
			}
		}
		return bestVal, bestState
	}
}

// func chooseMove3(s *implgame.State) implgame.State {
// 	m := TimedAlphaBeta(s, 3, s.NextPlayer == implgame.White, 500*1000*1000)
// 	return m
// }

func getMoveChooser(depth int, limit time.Duration) func(*implgame.State) implgame.State {
	return func(s *implgame.State) implgame.State {
		return TimedAlphaBeta(s, depth, s.NextPlayer == implgame.White, int64(limit))
	}
}

// func chooseMove2(s *implgame.State) implgame.State {
// 	m := TimedAlphaBeta(s, 2, s.NextPlayer == implgame.White, 500*1000*1000)
// 	return m
// }

func chooseMoveSimple(s *implgame.State) implgame.State {
	var chosenMove implgame.State
	var bestVal float64 = math.Inf(-1)
	futures := s.Futures()
	for _, f := range futures {
		testVal := myHeuristic(&f, s.NextPlayer)
		if testVal > bestVal {
			fmt.Printf("assigning chosenMove \n%v\n because %v >= %v\n", f, testVal, bestVal)
			bestVal = testVal
			chosenMove = f
		}
	}
	fmt.Printf("best move has heuristic of %v\n", bestVal)
	fmt.Printf("assigned move has heuristic of %v\n", myHeuristic(&chosenMove, s.NextPlayer))
	fmt.Printf("actual move: \n%v\n", chosenMove)
	return chosenMove
}
