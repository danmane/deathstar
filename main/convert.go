package main

import (
	"bytes"
	"encoding/json"
	"github.com/danmane/abalone/go/game"
	"github.com/danmane/sidious/implgame"
)

func inhale(s *game.State) *implgame.State {
	var ret *implgame.State
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&s); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(&buf).Decode(&ret); err != nil {
		panic(err)
	}
	return ret
}

func exhale(s *implgame.State) *game.State {
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
