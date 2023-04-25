package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    options := []string{"rock", "paper", "scissors"}

    // Generate a random integer between 0 and 2
    rand.Seed(time.Now().UnixNano())
    computerChoice := rand.Intn(3)

    // Read the player's choice 
    var playerChoice int
    fmt.Println("Enter 0 for rock, 1 for paper, 2 for scissors:")
    fmt.Scan(&playerChoice)

    // Print the choices
    fmt.Printf("Player choice: %s\n", options[playerChoice])
    fmt.Printf("Computer choice: %s\n", options[computerChoice])

    // Determine the winner
    if playerChoice == computerChoice {
        fmt.Println("It's a tie!")
    } else if (playerChoice == 0 && computerChoice == 2) ||
              (playerChoice == 1 && computerChoice == 0) ||
              (playerChoice == 2 && computerChoice == 1) {
        fmt.Println("Player wins!")
    } else {
        fmt.Println("Computer wins!")
    }
}

//return error if player picks too high or low of number