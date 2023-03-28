package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to Simon Says!")
	fmt.Println("Watch carefully and repeat the sequence.")

	sequence := generateSequence(5)
	fmt.Println("Simon says:", sequence)

	for i := 0; i < len(sequence); i++ {
		var input int
		fmt.Scanln(&input)

		if input != sequence[i] {
			fmt.Println("Game over!")
			return
		}
	}

	fmt.Println("Congratulations! You won!")
}

func generateSequence(length int) []int {
	sequence := make([]int, length)

	for i := 0; i < length; i++ {
		sequence[i] = rand.Intn(4) + 1
	}

	return sequence
}
