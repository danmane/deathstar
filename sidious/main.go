package main

import (
	"flag"
	"fmt"
	"github.com/danmane/abalone/go/api"
	"github.com/danmane/abalone/go/game"
	"github.com/danmane/abalone/go/quickstart"
	"time"
)

var (
	port      = flag.String("port", "3423", "port the ai runs on")
	depth     = flag.Int("depth", 3, "depth the search runs to")
	moves int = 0
)

func main() {
	flag.Parse()
	fmt.Println("Now, witness the firepower...")
	player := api.Player{Address: ":" + *port}
	quickstart.Play(player, wrapMove)
}

var defaultWeights = []int64{3000, 10, 2, 2, 5}

func wrapMove(s game.State, limit time.Duration) game.State {
	var mult int64
	mult = 1
	if s.NextPlayer == game.Black {
		mult = -1
	}
	inhaled := inhale(&s)
	currentH := calcHeuristic(inhaled, defaultWeights)
	fmt.Printf("Sidious (%v): %v\n", inhaled.NextPlayer.String(), currentH*mult)
	moveChooser := getMoveChooser(*depth, limit, defaultWeights)
	move := moveChooser(inhaled)
	fmt.Printf("value of heuristic after chosen move is %v\n", calcHeuristic(&move, defaultWeights)*mult)
	return *exhale(&move)
}
