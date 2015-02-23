package implgame

import (
	"github.com/danmane/abalone/go/game"
	"testing"
)

func Test_Inhale(t *testing.T) {
	standardIMPL := Standard
	standardConverted := Inhale(&game.Standard)
	if standardIMPL != *standardConverted {
		t.Error("conversion did not work", standardIMPL, *standardConverted)
	}
}

func Test_Exhale(t *testing.T) {
	standardIMPL := Standard
	thereAndBack := Inhale(Exhale(&standardIMPL))
	if standardIMPL != *thereAndBack {
		t.Error("conversion did not work", standardIMPL, *thereAndBack)
	}
}
