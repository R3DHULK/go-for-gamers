package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxAttempts = 3
)

type Team struct {
	Name    string
	Offense int
	Defense int
}

func (t *Team) OffenseChance() int {
	return rand.Intn(t.Offense) + 1
}

func (t *Team) DefenseChance() int {
	return rand.Intn(t.Defense) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Set up teams
	homeTeam := &Team{
		Name:    "Home Team",
		Offense: 80,
		Defense: 60,
	}
	awayTeam := &Team{
		Name:    "Away Team",
		Offense: 70,
		Defense: 50,
	}

	// Play game
	fmt.Printf("%s vs %s\n", homeTeam.Name, awayTeam.Name)
	homeScore, awayScore := 0, 0
	for i := 1; i <= 4; i++ {
		fmt.Printf("\nQuarter %d:\n", i)
		for j := 1; j <= MaxAttempts; j++ {
			// Home team's possession
			homeOffenseChance := homeTeam.OffenseChance()
			awayDefenseChance := awayTeam.DefenseChance()
			fmt.Printf("\nAttempt %d: %s offense (%d) vs %s defense (%d)\n", j, homeTeam.Name, homeOffenseChance, awayTeam.Name, awayDefenseChance)
			if homeOffenseChance > awayDefenseChance {
				fmt.Println("Touchdown!")
				homeScore += 7
				break
			} else {
				fmt.Println("No score.")
			}

			// Away team's possession
			awayOffenseChance := awayTeam.OffenseChance()
			homeDefenseChance := homeTeam.DefenseChance()
			fmt.Printf("\nAttempt %d: %s offense (%d) vs %s defense (%d)\n", j, awayTeam.Name, awayOffenseChance, homeTeam.Name, homeDefenseChance)
			if awayOffenseChance > homeDefenseChance {
				fmt.Println("Touchdown!")
				awayScore += 7
				break
			} else {
				fmt.Println("No score.")
			}
		}
	}

	// Game over
	fmt.Printf("\nFinal score: %s %d, %s %d\n", homeTeam.Name, homeScore, awayTeam.Name, awayScore)
	if homeScore > awayScore {
		fmt.Println("Home team wins!")
	} else if homeScore < awayScore {
		fmt.Println("Away team wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}
