package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    options := []string{"rock", "paper", "scissors", "knife"}

    rand.Seed(time.Now().UnixNano())
    computerChoice := RPS(rand.Intn(4))

    var playerChoice RPS
    fmt.Println("Enter 0 for rock, 1 for paper, 2 for scissors, 3 for knife:")
    fmt.Scan(&playerChoice)

    if playerChoice < 0 || playerChoice > 3 {
        fmt.Println("Invalid choice! Please pick a number between 0 and 3.")
        return
    }

    fmt.Printf("Player choice: %s\n", options[playerChoice])
    fmt.Printf("Computer choice: %s\n", options[computerChoice])

    if playerChoice == computerChoice {
        fmt.Println("It's a tie!")
    } else if (playerChoice == rock && computerChoice == scissors) ||
              (playerChoice == paper && computerChoice == rock) ||
              (playerChoice == scissors && computerChoice == paper) ||
              (playerChoice == knife && computerChoice == paper) ||
              (playerChoice == scissors && computerChoice == knife) {
        fmt.Println("Player wins!")
    } else {
        fmt.Println("Computer wins!")
    }
}
//enumeration
type RPS int 
const(
 rock RPS = iota
 paper
 scissors
 knife
)
