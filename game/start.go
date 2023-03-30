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
	if selectedGame == 1 {
		p := pokemon{}
	 	p.Play()
		
	} else if selectedGame == 2 {
		g := guessing{AttemptsAllowed: 3}
	 	g.Play()
	} else if selectedGame == 3 {
		r := speedRacer{}
	 	r.Play()
	} else if selectedGame == 4 {
		b := boardFace{}
	 	b.Play()
	}
	


}
