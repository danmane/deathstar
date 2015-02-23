package main

import (
	"fmt"
	"github.com/danmane/deathstar/implgame"
	"github.com/danmane/deathstar/sidious"
)

func playGame(state implgame.State, w, b sidious.HeuristicWeights, depth int) (score int) {
	for !state.GameOver() {
		var nextWeight sidious.HeuristicWeights
		if state.NextPlayer == implgame.White {
			nextWeight = w
		} else {
			nextWeight = b
		}
		state = sidious.AlphaBetaWrap(&state, sidious.WeightedHeuristic(nextWeight), depth)
	}
	o := state.Outcome()
	switch o {
	case implgame.WhiteWins:
		score = 1
	case implgame.BlackWins:
		score = -1
	case implgame.Tie:
		score = 0
	case implgame.NullOutcome:
		panic("got null outcome")
	}
	return
}

func playMatch(state implgame.State, pos, neg sidious.HeuristicWeights, depth int) int {
	return playGame(state, pos, neg, depth) - playGame(state, neg, pos, depth)
}

func main() {
	fmt.Println("not implemented")
}
