package game
import (
	"fmt"
	"math/rand"
	"time"
)

type Guessing struct {
	guess int 
	trialcount int 
}

func (g Guessing) Play(){
	g.trialcount = 1
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secrectNumber := randomizer.Intn(10)
	fmt.Println("Guess the number")
	for {
		fmt.Println("Please input your guess")
		fmt.Scan(&g.guess)
		if g.trialcount > 3 {
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

