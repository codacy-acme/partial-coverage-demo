package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func shuffledDeck() []Card {
	deck := NewDeck()
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func main() {
	deck := shuffledDeck()
	draw := func() Card {
		c := deck[0]
		deck = deck[1:]
		return c
	}

	player := []Card{draw(), draw()}
	dealer := []Card{draw(), draw()}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Blackjack! Dealer hits until 17.")

	for !IsBust(player) {
		fmt.Printf("Your hand: %s (%d)\n", RenderHand(player), HandValue(player))
		fmt.Printf("Dealer shows: %s\n", dealer[0])
		fmt.Print("Hit or stand? ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input:", err)
			return
		}
		choice := strings.TrimSpace(strings.ToLower(line))
		if choice == "stand" {
			break
		}
		if choice != "hit" {
			fmt.Println("Please type 'hit' or 'stand'.")
			continue
		}
		player = append(player, draw())
	}

	for DealerShouldHit(dealer) && !IsBust(player) {
		dealer = append(dealer, draw())
	}

	fmt.Printf("Your hand: %s (%d)\n", RenderHand(player), HandValue(player))
	fmt.Printf("Dealer hand: %s (%d)\n", RenderHand(dealer), HandValue(dealer))
	fmt.Println(DetermineWinner(player, dealer))
}
