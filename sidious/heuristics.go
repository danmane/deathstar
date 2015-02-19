package sidious

import (
	"github.com/danmane/sidious/implgame"
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

var Heuristics = []Heuristic{stones, segments, centrality, clusteredness, aggregateSegLengthSq}
