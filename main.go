package main

import (
	"fmt"
	"math/rand"
	"time"
)

var guess int

var trialcount int = 1

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secrectNumber := randomizer.Intn(10)

	fmt.Println("Guess the number")

	for {
		fmt.Println("Please input your guess")
		fmt.Scan(&guess)
		if trialcount > 3 {
			fmt.Println("You lose")
			break
		} else {
			if guess > secrectNumber {
				fmt.Println("Too Big")
			} else if guess < secrectNumber {
				fmt.Println("Too small")
			} else {
				fmt.Println("You win!")
				break
			}
		}
		trialcount++

	}
}
