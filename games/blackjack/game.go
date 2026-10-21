// Package main implements a simplified single-player Blackjack (vs. a
// dealer that hits until 17). Logic lives here, free of I/O, so it can be
// unit tested independently of the terminal front-end in main.go.
package main

import (
	"fmt"
	"strings"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var suits = []string{"♠", "♥", "♦", "♣"}

type Card struct {
	Rank string
	Suit string
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}

// Value returns the card's blackjack value, treating Ace as 1. Use
// HandValue for a full hand, which promotes a soft ace to 11 when it
// doesn't bust.
func (c Card) Value() int {
	return c.baseValue()
}

// baseValue returns a card's blackjack value, treating Ace as 1 (soft aces
// are promoted to 11 in HandValue when it doesn't bust the hand).
func (c Card) baseValue() int {
	switch c.Rank {
	case "J", "Q", "K":
		return 10
	case "A":
		return 1
	default:
		var v int
		fmt.Sscanf(c.Rank, "%d", &v)
		return v
	}
}

func NewDeck() []Card {
	deck := make([]Card, 0, len(ranks)*len(suits))
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}
	return deck
}

// HandValue returns the best blackjack total for a hand, counting at most
// one ace as 11 if doing so doesn't bust the hand.
func HandValue(hand []Card) int {
	total := 0
	aces := 0
	for _, c := range hand {
		total += c.baseValue()
		if c.Rank == "A" {
			aces++
		}
	}
	for aces > 0 && total+10 <= 21 {
		total += 10
		aces--
	}
	return total
}

func IsBust(hand []Card) bool {
	return HandValue(hand) > 21
}

type Result string

const (
	PlayerWins Result = "player wins"
	DealerWins Result = "dealer wins"
	PushResult Result = "push"
)

func DetermineWinner(playerHand, dealerHand []Card) Result {
	playerScore := HandValue(playerHand)
	dealerScore := HandValue(dealerHand)

	if IsBust(playerHand) {
		return DealerWins
	}
	if IsBust(dealerHand) {
		return PlayerWins
	}
	switch {
	case playerScore > dealerScore:
		return PlayerWins
	case dealerScore > playerScore:
		return DealerWins
	default:
		return PushResult
	}
}

// DealerShouldHit implements the standard fixed dealer strategy: hit on 16
// or below, stand on 17+.
func DealerShouldHit(dealerHand []Card) bool {
	return HandValue(dealerHand) < 17
}

// IsSoft17 reports whether a hand totals 17 by counting an ace as 11 (a
// "soft" 17, where some house rules make the dealer hit instead of stand).
func IsSoft17(hand []Card) bool {
	if HandValue(hand) != 17 {
		return false
	}
	hasAce := false
	total := 0
	for _, c := range hand {
		total += c.Value()
		if c.Rank == "A" {
			hasAce = true
		}
	}
	return hasAce && total+10 == 17
}

// IsBlackjack reports whether a 2-card starting hand is a natural 21.
func IsBlackjack(hand []Card) bool {
	return len(hand) == 2 && HandValue(hand) == 21
}

// RenderHand formats a hand as space-separated cards, e.g. "10♠ K♥".
func RenderHand(hand []Card) string {
	cards := make([]string, len(hand))
	for i, c := range hand {
		cards[i] = c.String()
	}
	return strings.Join(cards, " ")
}
