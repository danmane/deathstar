package main

import (
	"github.com/danmane/deathstar/sidious"
	"math/rand"
)

type Genetics sidious.HeuristicWeights

// general parameters for the genetic algorithm
type EvolutionControl struct {
	MutationProbability           float64
	InteractionsUntilReproduction int
}

// represents a population of HeuristicWeights
type Population struct {
	EC               EvolutionControl
	Size             int // num individuals in population
	Individuals      []Genetics
	Fitnesses        []int // fitness of each individual (by index). reset with each generation
	Interactions     int   // number of fitness-determining events that have occured thus far
	GenerationNumber int
}

func (p *Population) reproduce() {
	var rc RandomChooser
	for i, _ := range p.Individuals {
		rc.AddWeight(p.Fitnesses[i])
		p.Fitnesses[i] = 0
	}
	nextGeneration := make([]Genetics, p.Size)
	for i := 0; i < p.Size; i++ {
		p1 := p.Individuals[rc.Choose()]
		p2 := p.Individuals[rc.Choose()]
		child := getChild(p, &rc)
		child.mutate(p.EC)
		nextGeneration[i] = child
	}
	p.Individuals = nextGeneration
}

func getChild(p1, p2 *Genetics) (child Genetics) {
	length := len(p1)
	child = make(Genetics, length)
	for i := 0; i < length; i++ {
		if rand.Float32() < 0.5 {
			child[i] = (*p1)[i]
		} else {
			child[i] = (*p2)[i]
		}
	}
	return
}

func (g *Genetics) mutate(e *EvolutionControl) {
	for i := 0; i < len(*g); i++ {
		for rand.Float64() < e.MutationProbability {
			(*g)[i] ^= (1 << uint(rand.Intn(16)))
		}
	}
}
