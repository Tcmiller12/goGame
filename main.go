package main

import "github.com/temiller/goGame/game"


func main() {
	g := game.Guessing{NumberOfTriesAllowed:3}
	g.Play()

}

