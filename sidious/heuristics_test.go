package main

import (
	"github.com/danmane/deathstar/implgame"
	"testing"
)

func Benchmark_centrality(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		centrality(&implgame.Standard, implgame.White)
	}
}

func Benchmark_clusteredness(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		clusteredness(&implgame.Standard, implgame.White)
	}
}
