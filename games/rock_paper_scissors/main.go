package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rock, Paper, Scissors! Type rock, paper, or scissors.")
	fmt.Print("Your move: ")

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	player, err := ParseMove(strings.TrimSpace(strings.ToLower(line)))
	if err != nil {
		fmt.Println(err)
		return
	}

	opponent := RandomMove()
	fmt.Printf("Opponent played %s.\n", opponent)

	switch Play(player, opponent) {
	case Win:
		fmt.Println("You win!")
	case Lose:
		fmt.Println("You lose!")
	default:
		fmt.Println("It's a draw!")
	}
}
