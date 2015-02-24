package sidious

import (
	"github.com/danmane/deathstar/implgame"
	"math"
)

var NumHeuristics = 5
var PartialHeuristics = []PartialHeuristic{stones, segments, centrality, clusteredness, aggregateSegLengthSq}
var DefaultWeights = []int16{3000, 10, 2, 2, 5}

type PartialHeuristic func(g *implgame.State, p implgame.Player) int64
type Heuristic func(g *implgame.State) int64
type HeuristicWeights []int16

func stones(g *implgame.State, p implgame.Player) int64 {
	return int64(g.NumPieces(p))
}

func segments(g *implgame.State, p implgame.Player) int64 {
	return int64(len(g.Segments(p)))
}

func centrality(g *implgame.State, p implgame.Player) int64 {
	var aggregateDist int64 = 0
	pieces := g.Board.Pieces(p)
	for _, piece := range pieces.ToSlice() {
		aggregateDist += int64(piece.Dist2Origin())
	}
	return -aggregateDist
}

func clusteredness(g *implgame.State, p implgame.Player) int64 {
	var aggregateDist int64 = 0
	pieces := g.Board.Pieces(p)
	for _, piece1 := range pieces.ToSlice() {
		for _, piece2 := range pieces.ToSlice() {
			aggregateDist += int64(piece1.Dist2(&piece2))
		}
	}
	return -aggregateDist
}

func aggregateSegLengthSq(g *implgame.State, p implgame.Player) int64 {
	var aggregateLen int64 = 0
	segs := g.Segments(p)
	for _, s := range segs {
		aggregateLen += int64(s.Length * s.Length)
	}
	return aggregateLen
}

var DefaultHeuristic = WeightedHeuristic(DefaultWeights)

func WeightedHeuristic(weights HeuristicWeights) Heuristic {
	return func(state *implgame.State) int64 {

		if state.GameOver() {
			switch state.Outcome() {
			case implgame.WhiteWins:
				return math.MaxInt64
			case implgame.BlackWins:
				return math.MinInt64
			default:
				return 0.0
			}
		}
		var out int64 = 0
		for i, h := range PartialHeuristics {
			var val, weight int64
			val = h(state, implgame.White) - h(state, implgame.Black)
			weight = int64(weights[i])
			out += val * weight
		}
		return out
	}
}
