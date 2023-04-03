package game

import (
	"fmt"
	"math/rand"
	"time"
)

type pokemon struct {
	allCharacters []PokemonCharacter
	userSelection PokemonCharacter
}

func (p pokemon) getRandomStrength(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (p pokemon) start() {
	rand.Seed(time.Now().UnixNano())
	p.allCharacters = make([]PokemonCharacter, 6)
	names := []string{"Pikachu", "Blossomite", "Charmander", "Ditto", "Eevee", "Flareon"}
	// setup characters
	for i := 0; i < 6; i++ {
		p.allCharacters[i] = PokemonCharacter{Name: names[i], strength: p.getRandomStrength(80, 300)}
	}
	// ask user to pick one
	fmt.Println("Select the Pokemon Character You Want: ")
	fmt.Println("----------")
	fmt.Println("Characters:")
	fmt.Println("----------")
	for i, name := range names {
		fmt.Printf("     %d. %s \n", i+1, name)
	}
	selectedIndex := -1
	// get user selection from scan
	_, err := fmt.Scan(&selectedIndex)
	if err != nil {
		fmt.Println("The selection must be a number representing the character")
		return
	}
	if selectedIndex <= 0 || selectedIndex >= len(p.allCharacters) {
		fmt.Println("The selection is not valid")
		return
	}
	selectedIndex = selectedIndex - 1
	p.userSelection = p.allCharacters[selectedIndex]
	// select 3 random fighters
	// find 3 number between 0 and 5 that is not equal to the selected index and no duplicates
	number1 := getRandomNumber(5, selectedIndex)
	number2 := getRandomNumber(5, selectedIndex, number1)
	number3 := getRandomNumber(5, selectedIndex, number1, number2)

	fighter1 := p.allCharacters[number1]
	fighter2 := p.allCharacters[number2]
	fighter3 := p.allCharacters[number3]

	//Round1
	fmt.Printf("Round 1: Your Pokemon '%s' (strength = %d); Fights '%s' (strength = '%d')\n", p.userSelection.Name, p.userSelection.strength, fighter1.Name, fighter1.strength)
	winner, isDraw := p.fight(p.userSelection, fighter1)
	if winner != p.userSelection {
		fmt.Printf("Sorry you lost the fight to: %s", winner.Name)
		return
	}
	if !isDraw {
		fmt.Println("You won Round 1")
	}

	//Round2
	fmt.Printf("Round 2: Your Pokemon '%s' (strength = %d); Fights '%s' (strength = '%d')\n", p.userSelection.Name, p.userSelection.strength, fighter2.Name, fighter2.strength)
	winner, _ = p.fight(p.userSelection, fighter2)
	if winner != p.userSelection {
		fmt.Printf("Sorry you lost the fight to: %s", winner.Name)
		return
	}
	if !isDraw {
		fmt.Println("You won Round 2")
	}
	//Round3
	fmt.Printf("Round 3: Your Pokemon '%s' (strength = %d); Fights '%s' (strength = '%d')\n", p.userSelection.Name, p.userSelection.strength, fighter3.Name, fighter3.strength)
	winner, _ = p.fight(p.userSelection, fighter3)
	if winner != p.userSelection {
		fmt.Printf("Sorry you lost the fight to: %s", winner.Name)
		return
	}
	fmt.Println("Congratulations, you won all 3 rounds")
}

func (p pokemon) fight(fighter1, fighter2 PokemonCharacter) (PokemonCharacter, bool) {
	// edge case
	if fighter1.strength == fighter2.strength {
		return PokemonCharacter{Name: "unknown"}, true
	}
	if fighter1.strength > fighter2.strength {
		return fighter1, false
	}
	return fighter2, false
}

func (p pokemon) Play() {
	p.start()
}

type PokemonCharacter struct {
	Name     string
	strength int
}
