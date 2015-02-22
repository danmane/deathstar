package main

import (
	"github.com/danmane/deathstar/implgame"
	"testing"
)

func Benchmark_SortedFutures(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		SortedFutures(&implgame.Standard)
	}
}

func Benchmark_HeuristicFutures(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		futures := implgame.Standard.Futures()
		vals := make([]int64, len(futures))
		for i, f := range futures {
			vals[i] = calcHeuristic(&f, defaultWeights)
		}
	}
}

func Benchmark_FuturesWrapper(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		implgame.Standard.Futures()
	}
}
