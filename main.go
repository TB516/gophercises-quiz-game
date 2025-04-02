package main

//go:generate cp assets ./

import "github.com/TB516/gophercises-quiz-game/internal/game"

func main() {
	game.Run()
}
