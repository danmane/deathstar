package main

import (
	"fmt"
	"github.com/danmane/sidious/implgame"
	"math"
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

func chooseMove(s *implgame.State) implgame.State {
	_, m := Minimax(s, 2, s.NextPlayer == implgame.White)
	return m
}

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

func myHeuristic(state *implgame.State, whoami implgame.Player) float64 {
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
	weights := []float64{3000.0, 10.0, 2.0, 2.0, 5.0}
	var val float64 = 0
	for i, _ := range Heuristics {
		val += float64(Heuristics[i](state, whoami)) * weights[i]
		val -= float64(Heuristics[i](state, whoami.Next())) * weights[i]
	}
	return val
}
