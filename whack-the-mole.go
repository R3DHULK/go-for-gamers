package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NumHoles  = 6
	NumRounds = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Set up game
	score := 0
	for round := 1; round <= NumRounds; round++ {
		fmt.Printf("Round %d:\n", round)

		// Set up holes
		holes := make([]bool, NumHoles)
		for i := 0; i < NumHoles; i++ {
			holes[i] = false
		}

		// Randomly select a hole for the mole
		moleIndex := rand.Intn(NumHoles)
		holes[moleIndex] = true

		// Play round
		for i, hole := range holes {
			if hole {
				fmt.Printf("[%d]  M  ", i+1)
			} else {
				fmt.Printf("[%d]     ", i+1)
			}
		}
		fmt.Println()

		var input int
		fmt.Print("Enter the number of the hole to whack: ")
		fmt.Scanln(&input)

		if input < 1 || input > NumHoles {
			fmt.Println("Invalid input.")
			continue
		}

		if holes[input-1] {
			score += 1
			fmt.Println("Whacked the mole!")
		} else {
			fmt.Println("Missed the mole!")
		}
	}

	// Game over
	fmt.Printf("Game over! Final score: %d\n", score)
}
