package main

import "github.com/temiller/goGame/game"

func main() {
	g := game.Guessing{AttemptsAllowed: 3}
	g.Play()

}
