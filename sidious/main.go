package sidious

import (
	"flag"
	"time"

	"github.com/danmane/abalone/go/api"
	"github.com/danmane/abalone/go/game"
	"github.com/danmane/abalone/go/quickstart"
)

func bustaMove(s game.State, limit time.Duration) game.State {
	maximize := s.NextPlayer == game.White
	_, move := Minimax(inhale(&s), 3, maximize)
	return *exhale(move)
}

var (
	port = flag.String("port", "3423", "port the ai runs on")
)

func main() {
	flag.Parse()
	player := api.Player{Address: ":" + *port}
	quickstart.Play(player, bustaMove)
}
