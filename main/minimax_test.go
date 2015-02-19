package main

import (
	"github.com/danmane/sidious/implgame"
	"testing"
)

var s implgame.State = implgame.Standard

func Benchmark_bustaMove_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&s, 1, true, implgame.White)
	}
}

func Benchmark_bustaMove_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&s, 2, true, implgame.White)
	}
}

// func Benchmark_bustaMove_3(b *testing.B) {
// 	for n := 0; n <= b.N; n++ {
// 		Minimax(&s, 3)
// 	}
// }
