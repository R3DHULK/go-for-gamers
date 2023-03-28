package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Clue struct {
	Description string
	IsSolution  bool
}

type Case struct {
	Title    string
	Clues    []Clue
	Solution string
}

func (c *Case) Investigate() {
	fmt.Println(c.Title)
	fmt.Println("-------------------------------")
	for _, clue := range c.Clues {
		fmt.Println(clue.Description)
	}
}

func (c *Case) Solve(guess string) bool {
	return strings.ToLower(guess) == strings.ToLower(c.Solution)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Detective Game!")
	fmt.Println("You are a detective trying to solve a case.")

	case1 := Case{
		Title: "Case 1",
		Clues: []Clue{
			{Description: "The victim's neighbor heard a loud argument the night of the murder.", IsSolution: false},
			{Description: "The victim was found with a knife in their back.", IsSolution: false},
			{Description: "The victim's ex-partner had a motive for the murder.", IsSolution: true},
		},
		Solution: "ex-partner",
	}

	case2 := Case{
		Title: "Case 2",
		Clues: []Clue{
			{Description: "The victim was last seen at a bar on the night of the murder.", IsSolution: false},
			{Description: "The victim's body was found in an alleyway.", IsSolution: false},
			{Description: "A witness reported seeing a man in a red hoodie running from the scene of the crime.", IsSolution: true},
		},
		Solution: "man in red hoodie",
	}

	cases := []Case{case1, case2}

	for _, c := range cases {
		c.Investigate()

		var guess string
		fmt.Print("Who do you think is the culprit? ")
		guess, _ = reader.ReadString('\n')
		guess = strings.TrimSpace(guess)

		if c.Solve(guess) {
			fmt.Println("Congratulations! You solved the case.")
		} else {
			fmt.Println("Sorry, that's not the correct answer.")
		}
	}
}
