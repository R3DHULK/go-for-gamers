package main

import (
    "bufio"
    "fmt"
    "os"
    "math/rand"
    "strings"
)

type question struct {
    text string
    choices []string
    answer int
}

func main() {
    // Define the questions for the trivia game
    questions := []question{
        {
            text: "What is the capital of France?",
            choices: []string{"A) Paris", "B) Rome", "C) Berlin", "D) Madrid"},
            answer: 0,
        },
        {
            text: "Who wrote the novel 'To Kill a Mockingbird'?",
            choices: []string{"A) Harper Lee", "B) J.K. Rowling", "C) Charles Dickens", "D) William Shakespeare"},
            answer: 0,
        },
        {
            text: "What is the smallest planet in our solar system?",
            choices: []string{"A) Mercury", "B) Venus", "C) Mars", "D) Pluto"},
            answer: 0,
        },
    }

    // Shuffle the order of the questions
    shuffled := shuffle(questions)

    // Initialize the score
    score := 0

    // Loop through each question and ask the player
    for _, q := range shuffled {
        fmt.Println(q.text)
        for _, c := range q.choices {
            fmt.Println(c)
        }
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Your answer: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == string(q.answer + 1 + 'A') {
            fmt.Println("Correct!")
            score++
        } else {
            fmt.Println("Incorrect.")
        }
    }

    // Print the final score
    fmt.Printf("You scored %d out of %d.\n", score, len(questions))
}

// shuffle takes a slice of questions and returns a new slice with the order shuffled
func shuffle(questions []question) []question {
    shuffled := make([]question, len(questions))
    copy(shuffled, questions)
    for i := range shuffled {
        j := rand.Intn(i + 1)
        shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
    }
    return shuffled
}
