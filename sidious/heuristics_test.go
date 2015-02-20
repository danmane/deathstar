package main

import (
	"github.com/danmane/deathstar/implgame"
	"testing"
)

func Benchmark_Heur_centrality(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		centrality(&implgame.Standard, implgame.White)
	}
}

func Benchmark_Heur_clusteredness(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		clusteredness(&implgame.Standard, implgame.White)
	}
}

func Benchmark_Heur_aggregateSegLengthSq(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		aggregateSegLengthSq(&implgame.Standard, implgame.White)
	}
}

func Benchmark_myHeuristic(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		myHeuristic(&implgame.Standard, implgame.White)
	}
}
