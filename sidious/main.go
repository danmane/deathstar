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

// func inhale(s *game.State) *implgame.State {
// 	return &implgame.State{
// 		Board:          inhaleB(&s.Board),
// 		NextPlayer:     implgame.Player(s.NextPlayer),
// 		MovesRemaining: s.MovesRemaining,
// 		MarblesPerMove: s.MarblesPerMove,
// 		LossThreshold:  s.LossThreshold,
// 	}
// }
// func exhale(s *implgame.State) *game.State {
// 	return &game.State{
// 		Board:          exhaleB(&s.Board),
// 		NextPlayer:     game.Player(s.NextPlayer),
// 		MovesRemaining: s.MovesRemaining,
// 		MarblesPerMove: s.MarblesPerMove,
// 		LossThreshold:  s.LossThreshold,
// 	}
// }

// func exhaleB(b *implgame.Board) game.Board {
// 	return game.Board{
// 		WhitePositions: game.HexSet(b.WhitePositions),
// 		BlackPositions: game.HexSet(b.BlackPositions),
// 		EdgeLength:     b.EdgeLength,
// 	}
// }
// func inhaleB(b *game.Board) implgame.Board {
// 	return implgame.Board{
// 		WhitePositions: implgame.HexSet(b.WhitePositions),
// 		BlackPositions: implgame.HexSet(b.BlackPositions),
// 		EdgeLength:     b.EdgeLength,
// 	}
// }

// func inhaleHS (hs game.HexSet) implgame.HexSet {

// }

var (
	port = flag.String("port", "3423", "port the ai runs on")
)

func main() {
	flag.Parse()
	player := api.Player{Address: ":" + *port}
	quickstart.Play(player, bustaMove)
}
