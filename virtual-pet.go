package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pet struct {
	Name   string
	Hunger int
	Mood   string
}

func (p *Pet) Feed() {
	p.Hunger -= 10
}

func (p *Pet) Play() {
	if p.Mood == "happy" {
		p.Hunger += 10
	} else {
		p.Hunger += 20
	}
	p.Mood = "happy"
}

func (p *Pet) Tick() {
	p.Hunger += 5
	if p.Hunger > 100 {
		p.Hunger = 100
		p.Mood = "angry"
	}
	if p.Hunger > 50 && p.Mood == "happy" {
		p.Mood = "hungry"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Virtual Pet!")

	fmt.Print("What's your pet's name? ")
	name, _ := reader.ReadString('\n')

	pet := &Pet{
		Name:   name,
		Hunger: 50,
		Mood:   "happy",
	}

	fmt.Println("Your pet has been created!")
	fmt.Println("Type 'feed' to feed your pet")
	fmt.Println("Type 'play' to play with your pet")
	fmt.Println("Type 'exit' to quit the game")

	for {
		fmt.Printf("Hunger: %d, Mood: %s\n", pet.Hunger, pet.Mood)

		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = command[:len(command)-1]

		switch command {
		case "feed":
			pet.Feed()
			fmt.Println("You fed your pet!")
		case "play":
			pet.Play()
			fmt.Println("You played with your pet!")
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command")
		}

		pet.Tick()
	}
}
