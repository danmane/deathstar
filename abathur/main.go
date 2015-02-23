package main

import (
	"fmt"
	"github.com/danmane/deathstar/implgame"
	"github.com/danmane/deathstar/sidious"
)

func playGame(w, b sidious.HeuristicWeights, depth int) (score int) {
	state := implgame.Standard
	for !state.GameOver() {
		fmt.Printf("move %v, wPieces %v blackPieces %v\n", state.MovesRemaining, state.NumPieces(implgame.White), state.NumPieces(implgame.Black))
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

func playMatch(pos, neg sidious.HeuristicWeights, depth int) int {
	return playGame(pos, neg, depth) - playGame(neg, pos, depth)
}

func main() {
	fmt.Println("not implemented")
}
