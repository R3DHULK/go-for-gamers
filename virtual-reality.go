package main

import (
	"fmt"
	"time"
)

type Player struct {
	X, Y   int
	Health int
}

func main() {
	fmt.Println("Welcome to Virtual Reality Game!")

	player := Player{
		X:      5,
		Y:      5,
		Health: 100,
	}

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Draw the game world
		fmt.Println("###########################################")
		for i := 0; i < 20; i++ {
			fmt.Print("#")
			for j := 0; j < 38; j++ {
				if i == player.Y && j == player.X {
					fmt.Print("O")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println("#")
		}
		fmt.Println("###########################################")

		// Check for player input
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)

		// Update player position based on input
		switch input {
		case "w":
			player.Y -= 1
		case "a":
			player.X -= 1
		case "s":
			player.Y += 1
		case "d":
			player.X += 1
		}

		// Update player health over time
		player.Health -= 1

		// End the game if the player runs out of health
		if player.Health <= 0 {
			fmt.Println("Game over!")
			return
		}

		// Wait for a moment before redrawing the screen
		time.Sleep(50 * time.Millisecond)
	}
}
