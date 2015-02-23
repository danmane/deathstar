package sidious

import (
	"github.com/danmane/deathstar/implgame"
	"testing"
)

func Benchmark_Minimax_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&implgame.Standard, 1, true, DefaultWeights)
	}
}

func Benchmark_Minimax_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&implgame.Standard, 2, true, DefaultWeights)
	}
}
