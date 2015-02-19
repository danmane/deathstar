package main

import (
	"github.com/danmane/abalone/go/game"
	"github.com/danmane/deathstar/implgame"
	"testing"
)

func Test_inhale(t *testing.T) {
	standardIMPL := implgame.Standard
	standardConverted := inhale(&game.Standard)
	if !standardIMPL.Eq(standardConverted) {
		t.Error("conversion did not work")
	}
}

func Test_exhale(t *testing.T) {
	standardIMPL := implgame.Standard
	thereAndBack := inhale(exhale(&standardIMPL))
	if !standardIMPL.Eq(thereAndBack) {
		t.Error("conversion did not work")
	}
}
