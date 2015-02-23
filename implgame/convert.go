package implgame

import (
	"bytes"
	"encoding/json"
	"github.com/danmane/abalone/go/game"
)

func Inhale(s *game.State) *State {
	var ret *State
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&s); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(&buf).Decode(&ret); err != nil {
		panic(err)
	}
	return ret
}

func Exhale(s *State) *game.State {
	var ret *game.State
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&s); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(&buf).Decode(&ret); err != nil {
		panic(err)
	}
	return ret
}
