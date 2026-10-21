package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	const min, max, maxAttempts = 1, 100, 10
	secret := min + rand.Intn(max-min+1)
	game, err := NewGame(secret, min, max, maxAttempts)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Guess the Number! I'm thinking of a number between %d and %d.\n", min, max)

	for !game.IsOver() {
		fmt.Printf("Attempts remaining: %d\nYour guess: ", game.AttemptsRemaining())
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input:", err)
			return
		}
		guess, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Please enter a whole number.")
			continue
		}

		switch game.Guess(guess) {
		case Correct:
			fmt.Printf("Correct! The number was %d.\n", secret)
			return
		case Higher:
			fmt.Println("Higher!")
		case Lower:
			fmt.Println("Lower!")
		}
	}

	fmt.Printf("Out of attempts! The number was %d.\n", secret)
}
