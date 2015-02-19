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
	moves int = 0
)

func main() {
	flag.Parse()
	fmt.Println("Now, witness the firepower...")
	player := api.Player{Address: ":" + *port}
	quickstart.Play(player, wrapMove)
}

func wrapMove(s game.State, limit time.Duration) game.State {
	inhaled := inhale(&s)
	currentH := myHeuristic(inhaled, inhaled.NextPlayer)
	fmt.Printf("i believe I am %v\n", inhaled.NextPlayer.String())
	fmt.Printf("rec game in which my heuristic is %v\n", currentH)
	move := chooseMove(inhaled)
	fmt.Printf("value of heuristic after chosen move is %v\n", myHeuristic(&move, inhaled.NextPlayer))
	return *exhale(&move)
}
