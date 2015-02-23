package main

import (
	"github.com/danmane/deathstar/implgame"
	"testing"
)

// var saneWeights, suicideWeights sidious.HeuristicWeights
var saneWeights = []int64{1000, 1, 1, 1, 1}
var suicideWeights = []int64{-100, 0, -1, 0, 0}

func Test_playGame(t *testing.T) {
	fastGame := implgame.Standard
	fastGame.MovesRemaining = 5
	if playGame(fastGame, saneWeights, suicideWeights, 1) != 1 {
		t.Error("sanity did not beat suicidal")
	}
	if playGame(fastGame, suicideWeights, saneWeights, 1) != -1 {
		t.Error("suicidal did not lose to sanity")
	}
}

func Test_playMatch(t *testing.T) {
	fastGame := implgame.Standard
	fastGame.MovesRemaining = 5

	if playMatch(fastGame, saneWeights, suicideWeights, 1) != 2 {
		t.Error("sanity did not beat suicidal (match)")
	}
	if playMatch(fastGame, suicideWeights, saneWeights, 1) != -2 {
		t.Error("suicidal did not lose to sanity (match)")
	}
}
