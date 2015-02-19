package main

import (
	"github.com/danmane/sidious/implgame"
	"math"
	"testing"
)

var s implgame.State = implgame.Standard

func Benchmark_Minimax_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&s, 1, true)
	}
}

func Benchmark_Minimax_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&s, 2, true)
	}
}

func Benchmark_AlphaBeta_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 1, math.Inf(-1), math.Inf(1), true)
	}
}

func Benchmark_AlphaBeta_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 2, math.Inf(-1), math.Inf(1), true)
	}
}

func Benchmark_AlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 3, math.Inf(-1), math.Inf(1), true)
	}
}

func Benchmark_TimedAlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		TimedAlphaBeta(&s, 3, true, 3*1000*1000*1000)
	}
}

// func Benchmark_bustaMove_3(b *testing.B) {
// 	for n := 0; n <= b.N; n++ {
// 		Minimax(&s, 3)
// 	}
// }
