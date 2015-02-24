package main

import (
	"math/rand"
	"sort"
)

type RandomChooser struct {
	totalFitness      int
	cumulativeFitness []int
}

func (rc *RandomChooser) AddWeight(fitness int) {
	rc.totalFitness += fitness
	rc.cumulativeFitness = append(rc.cumulativeFitness, rc.totalFitness)
}

func (rc *RandomChooser) choose(x int) int {
	if x < 0 || x >= rc.totalFitness {
		panic("invalid index")
	}
	return sort.SearchInts(rc.cumulativeFitness, x+1)
}

func (rc *RandomChooser) Choose() int {
	return rc.choose(rand.Intn(rc.totalFitness))
}
