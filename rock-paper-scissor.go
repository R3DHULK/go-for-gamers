package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Set up game
	rand.Seed(time.Now().UnixNano())
	choices := []string{"rock", "paper", "scissors"}
	wins := [][]int{
		{0, 1, -1},
		{-1, 0, 1},
		{1, -1, 0},
	}

	// Play game
	fmt.Println("Let's play Rock-Paper-Scissors!")
	for {
		// Get player choice
		fmt.Print("Enter your choice (rock/paper/scissors): ")
		playerChoice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		playerChoice = strings.TrimSpace(playerChoice)

		// Check player choice is valid
		playerIndex := -1
		for i, choice := range choices {
			if choice == playerChoice {
				playerIndex = i
				break
			}
		}
		if playerIndex == -1 {
			fmt.Println("Invalid choice, please enter rock, paper, or scissors.")
			continue
		}

		// Get computer choice
		computerIndex := rand.Intn(3)
		fmt.Println("Computer chose", choices[computerIndex])

		// Determine winner
		result := wins[playerIndex][computerIndex]
		switch result {
		case 1:
			fmt.Println("You win!")
		case 0:
			fmt.Println("It's a tie!")
		case -1:
			fmt.Println("Computer wins!")
		}

		// Ask if player wants to play again
		fmt.Print("Play again? (y/n): ")
		playAgain, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		playAgain = strings.TrimSpace(playAgain)
		if playAgain != "y" {
			break
		}
	}
}
