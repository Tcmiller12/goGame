package game

import "fmt"

type Boardface struct {
}

func (b *Boardface) Play() {
	fmt.Print("Sorry, this Game has not being implemented. Come back again")
}