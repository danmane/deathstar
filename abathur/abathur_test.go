package main

import (
	"github.com/danmane/deathstar/implgame"
	"testing"
)

// var saneWeights, suicideWeights sidious.HeuristicWeights
var saneWeights = []int64{1, 0, 0, 0, 0}
var suicideWeights = []int64{-1, 0, 0, 0, 0}

func Test_playGame(t *testing.T) {
	if playGame(saneWeights, suicideWeights, 1) != 1 {
		t.Error("sanity did not beat suicidal")
	}
	if playGame(suicideWeights, saneWeights, 1) != -1 {
		t.Error("suicidal did not lose to sanity")
	}
}

func Test_playMatch(t *testing.T) {
	if playMatch(saneWeights, suicideWeights, 1) != 2 {
		t.Error("sanity did not beat suicidal (match)")
	}
	if playMatch(suicideWeights, saneWeights, 1) != -2 {
		t.Error("suicidal did not lose to sanity (match)")
	}
}
