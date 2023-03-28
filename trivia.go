package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
    "io"
)

func main() {
	// Load questions from CSV file
	questions, err := loadQuestions("questions.csv")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		os.Exit(1)
	}

	// Shuffle the questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	// Start the game
	reader := bufio.NewReader(os.Stdin)
	correctAnswers := 0
	totalQuestions := len(questions)
	fmt.Printf("Welcome to the trivia game! You will be asked %d questions.\n", totalQuestions)
	fmt.Println("Press enter to start...")
	reader.ReadString('\n')

	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q[0])
		answer := strings.TrimSpace(q[1])

		// Prompt the user for an answer
		fmt.Print("Your answer: ")
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)

		// Check if the answer is correct
		if strings.EqualFold(userAnswer, answer) {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Printf("Incorrect. The correct answer is %s.\n", answer)
		}

		// Pause for a moment before displaying the next question
		time.Sleep(1 * time.Second)
	}

	// Display the final score
	fmt.Printf("Game over! You got %d out of %d questions correct.\n", correctAnswers, totalQuestions)
}

func loadQuestions(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var questions [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		questions = append(questions, line)
	}

	return questions, nil
}
