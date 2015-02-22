package main

import (
	"github.com/danmane/abalone/go/game"
	"github.com/danmane/deathstar/implgame"
	"testing"
)

func Test_inhale(t *testing.T) {
	standardIMPL := implgame.Standard
	standardConverted := inhale(&game.Standard)
	if standardIMPL != *standardConverted {
		t.Error("conversion did not work", standardIMPL, *standardConverted)
	}
}

func Test_exhale(t *testing.T) {
	standardIMPL := implgame.Standard
	thereAndBack := inhale(exhale(&standardIMPL))
	if standardIMPL != *thereAndBack {
		t.Error("conversion did not work", standardIMPL, *thereAndBack)
	}
}
