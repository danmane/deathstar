package sidious

import (
	"github.com/danmane/abalone/go/game"
	"github.com/danmane/sidious/implgame"
	"testing"
)

func Test_inhale(t *testing.T) {
	standardIMPL := implgame.Standard
	standardConverted := inhale(&game.Standard)
	if !standardIMPL.Eq(standardConverted) {
		t.Error("conversion did not work")
	}
}
