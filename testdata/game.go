package main

import (
	"encoding/json"
	"fmt"
)

// This is declared ahead of Game and the constants so that we ensure that even
// if there are parsing errors (e.g. undefined reference to Game_Parse), the
// enums will be generated.
func ck(game Game, str string) {
	if fmt.Sprint(game) != str {
		panic("game.go.string: " + str)
	}
	if v, err := Game_Parse(str); err != nil {
		panic("game.go.parse: " + str + ": ERR:" + err.Error())
	} else if v != game {
		panic("game.go.parse: " + str)
	}
	if s, err := game.MarshalText(); err != nil {
		panic("game.go.marshal: " + str + ": ERR:" + err.Error())
	} else if string(s) != str {
		panic("game.go.marshal: " + str)
	}
	var g Game
	if err := g.UnmarshalText([]byte(str)); err != nil {
		panic("game.go.unmarshal: " + str + ": ERR:" + err.Error())
	} else if g != game {
		panic("game.go.unmarshal: " + str)
	}

	if data, err := json.Marshal(game); err != nil {
		panic("game.go.json: " + str + ": ERR: " + err.Error())
	} else if string(data) != fmt.Sprintf("%q", str) {
		panic("game.go.json: " + str + ": got: " + string(data))
	}
}

type Game int

const (
	CHESS       Game = iota // "Chess"
	GO                      // "The other Go"
	DEFAULT                 // This comment doesn't affect the name
	CHECKERS                // "The jumpy game on a "chess" board"
	TIDDLYWINKS             // "The jumpy game NOT on a chess board"
)

func main() {
	ck(DEFAULT, "DEFAULT")
	ck(CHESS, "Chess")
	ck(GO, "The other Go")
	ck(CHECKERS, `The jumpy game on a "chess" board`)
}
