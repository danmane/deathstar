package main

import (
	"github.com/danmane/deathstar/implgame"
	"math"
)

type Heuristic func(g *implgame.State, p implgame.Player) int

func stones(g *implgame.State, p implgame.Player) int {
	return g.NumPieces(p)
}

func segments(g *implgame.State, p implgame.Player) int {
	return len(g.Segments(p))
}

func centrality(g *implgame.State, p implgame.Player) int {
	var aggregateDist int = 0
	pieces := g.Board.Pieces(p)
	for piece, _ := range pieces {
		aggregateDist += piece.Dist2Origin()
	}
	return -aggregateDist
}

func clusteredness(g *implgame.State, p implgame.Player) int {
	var aggregateDist int = 0
	pieces := g.Board.Pieces(p)
	for piece1, _ := range pieces {
		for piece2, _ := range pieces {
			aggregateDist += piece1.Dist2(&piece2)
		}
	}
	return -aggregateDist
}

func aggregateSegLengthSq(g *implgame.State, p implgame.Player) int {
	var aggregateLen int = 0
	segs := g.Segments(p)
	for _, s := range segs {
		aggregateLen += s.Length * s.Length
	}
	return aggregateLen
}

func myHeuristic(state *implgame.State, whoami implgame.Player) float64 {
	if state.GameOver() {
		switch state.Outcome() {
		case implgame.WhiteWins:
			return math.Inf(1)
		case implgame.BlackWins:
			return math.Inf(-1)
		default:
			return 0.0
		}
	}
	weights := []float64{3000.0, 10.0, 2.0, 2.0, 5.0}
	var val float64 = 0
	for i, _ := range Heuristics {
		val += float64(Heuristics[i](state, whoami)) * weights[i]
		val -= float64(Heuristics[i](state, whoami.Next())) * weights[i]
	}
	return val
}

var Heuristics = []Heuristic{stones, segments, centrality, clusteredness, aggregateSegLengthSq}
