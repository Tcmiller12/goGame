package game

import (
	"fmt"
)

type guessing struct {
	guess           int
	trialcount      int
	AttemptsAllowed int
}

func (g guessing) Play() {
	g.trialcount = 1
	secrectNumber := getRandomNumber(10)
	fmt.Println("Guess the number")
	for {
		fmt.Println("Please input your guess")
		fmt.Scan(&g.guess)
		if g.trialcount > g.AttemptsAllowed {
			fmt.Println("You lose")
			break
		} else {
			if g.guess > secrectNumber {
				fmt.Println("Too Big")
			} else if g.guess < secrectNumber {
				fmt.Println("Too small")
			} else {
				fmt.Println("You win!")
				break
			}
		}
		g.trialcount++

	}
}
