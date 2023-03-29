package game

import "fmt"

func Start() {
	var selectedGame int
	// display game menu
	menu := `
	What game would you like to play today?
		1. Pokemon
		2. Guessing
		3. SpeedRacer
		4. Boardface
	`
	fmt.Println(menu)
	fmt.Scan(&selectedGame)
	switch selectedGame {
	case 1:
		p := Pokemon{}
		p.Play()
	case 2:
		g := Guessing{AttemptsAllowed: 3}
		g.Play()
	case 3:
		r := Racer{}
		r.Play()
	case 4:
		b := Boardface{}
		b.Play()
	default:
		fmt.Print("Option selected is not valid")
	}

}