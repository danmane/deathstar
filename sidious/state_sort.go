package sidious

import (
	"github.com/danmane/deathstar/implgame"
	"sort"
)

type SortedStates struct {
	States []implgame.State
	Vals   []int64
}

func (s SortedStates) Len() int {
	return len(s.States)
}
func (s SortedStates) Swap(i, j int) {
	s.States[i], s.States[j] = s.States[j], s.States[i]
	s.Vals[i], s.Vals[j] = s.Vals[j], s.Vals[i]
}
func (s SortedStates) Less(i, j int) bool {
	return s.Vals[i] < s.Vals[i]
}

func SortedFutures(s *implgame.State) SortedStates {
	futures := s.Futures()
	vals := make([]int64, len(futures))
	for i, f := range futures {
		vals[i] = DefaultHeuristic(&f)
	}
	sortable := SortedStates{States: futures, Vals: vals}
	sort.Sort(sortable)
	return sortable
}
