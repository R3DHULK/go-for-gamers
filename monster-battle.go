package main

import (
	"fmt"
	"math/rand"
	"time"
)

type monster struct {
	name   string
	health int
	attack int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	playerMonster := monster{"Player Monster", 100, 20}
	enemyMonster := monster{"Enemy Monster", 100, 20}

	fmt.Println("A wild", enemyMonster.name, "appeared!")

	for playerMonster.health > 0 && enemyMonster.health > 0 {
		fmt.Println("\nPlayer Monster: ", playerMonster)
		fmt.Println("Enemy Monster: ", enemyMonster)

		fmt.Println("\nWhat will you do?")
		fmt.Println("1. Attack")
		fmt.Println("2. Heal")
		fmt.Print("Choose an action: ")
		var action int
		fmt.Scanln(&action)

		switch action {
		case 1:
			attack(&playerMonster, &enemyMonster)
		case 2:
			heal(&playerMonster)
		default:
			fmt.Println("Invalid action. Please try again.")
		}

		if enemyMonster.health > 0 {
			attack(&enemyMonster, &playerMonster)
		}
	}

	if playerMonster.health > 0 {
		fmt.Println("\nCongratulations! You defeated the enemy", enemyMonster.name)
	} else {
		fmt.Println("\nGame over. The enemy", enemyMonster.name, "defeated your monster.")
	}
}

func attack(attacker *monster, defender *monster) {
	damage := attacker.attack + rand.Intn(10)
	defender.health -= damage

	fmt.Println(attacker.name, "attacked", defender.name, "for", damage, "damage!")
}

func heal(monster *monster) {
	healAmount := 30 + rand.Intn(20)
	monster.health += healAmount

	if monster.health > 100 {
		monster.health = 100
	}

	fmt.Println(monster.name, "healed for", healAmount, "HP!")
}
