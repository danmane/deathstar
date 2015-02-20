package main

import (
	"github.com/danmane/deathstar/implgame"
	"sort"
)

type SortedStates struct {
	states []implgame.State
	vals   []float64
}

func (s SortedStates) Len() int {
	return len(s.states)
}
func (s SortedStates) Swap(i, j int) {
	s.states[i], s.states[j] = s.states[j], s.states[i]
	s.vals[i], s.vals[j] = s.vals[j], s.vals[i]
}
func (s SortedStates) Less(i, j int) bool {
	return s.vals[i] < s.vals[i]
}

func SortedFutures(s *implgame.State) SortedStates {
	futures := s.Futures()
	vals := make([]float64, len(futures))
	for i, f := range futures {
		vals[i] = myHeuristic(&f, implgame.White)
	}
	sortable := SortedStates{states: futures, vals: vals}
	sort.Sort(sortable)
	return sortable
}
