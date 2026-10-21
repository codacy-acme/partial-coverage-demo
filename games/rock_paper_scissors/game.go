// Package main implements Rock-Paper-Scissors. Logic lives here, free of I/O,
// so it can be unit tested independently of the terminal front-end in main.go.
package main

import (
	"fmt"
	"math/rand"
)

type Move string

const (
	Rock     Move = "rock"
	Paper    Move = "paper"
	Scissors Move = "scissors"
)

var AllMoves = []Move{Rock, Paper, Scissors}

var validMoves = map[Move]bool{Rock: true, Paper: true, Scissors: true}

// beats reports whether move a defeats move b.
var beats = map[Move]Move{
	Rock:     Scissors,
	Scissors: Paper,
	Paper:    Rock,
}

func ParseMove(s string) (Move, error) {
	m := Move(s)
	if !validMoves[m] {
		return "", fmt.Errorf("invalid move %q: must be rock, paper, or scissors", s)
	}
	return m, nil
}

// Outcome is from the perspective of the player (first argument).
type Outcome string

const (
	Win  Outcome = "win"
	Lose Outcome = "lose"
	Draw Outcome = "draw"
)

func Play(player, opponent Move) Outcome {
	if player == opponent {
		return Draw
	}
	if beats[player] == opponent {
		return Win
	}
	return Lose
}

// Opposite flips an outcome to the other player's perspective: a Win for
// one side is a Lose for the other, and a Draw stays a Draw.
func Opposite(outcome Outcome) Outcome {
	switch outcome {
	case Win:
		return Lose
	case Lose:
		return Win
	default:
		return Draw
	}
}

// Beats returns the move that the given move defeats, e.g. Rock -> Scissors.
func Beats(move Move) Move {
	return beats[move]
}

// RandomMove picks a uniformly random move for a computer opponent.
func RandomMove() Move {
	return AllMoves[rand.Intn(len(AllMoves))]
}
