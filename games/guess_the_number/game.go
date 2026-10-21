// Package main implements Guess-the-Number. Logic lives here, free of I/O,
// so it can be unit tested independently of the terminal front-end in main.go.
package main

import "fmt"

type Hint string

const (
	Correct Hint = "correct"
	Higher  Hint = "higher" // the secret is higher than the guess
	Lower   Hint = "lower"  // the secret is lower than the guess
)

type Game struct {
	Secret      int
	Min, Max    int
	Attempts    int
	MaxAttempts int
	History     []int
}

func NewGame(secret, min, max, maxAttempts int) (*Game, error) {
	if secret < min || secret > max {
		return nil, fmt.Errorf("secret %d must be within [%d, %d]", secret, min, max)
	}
	return &Game{Secret: secret, Min: min, Max: max, MaxAttempts: maxAttempts}, nil
}

func (g *Game) Guess(n int) Hint {
	g.Attempts++
	g.History = append(g.History, n)
	switch {
	case n == g.Secret:
		return Correct
	case n < g.Secret:
		return Higher
	default:
		return Lower
	}
}

func (g *Game) IsOver() bool {
	return g.Attempts >= g.MaxAttempts
}

func (g *Game) AttemptsRemaining() int {
	return g.MaxAttempts - g.Attempts
}

// LastGuess returns the most recent guess, or ok=false if no guess has
// been made yet.
func (g *Game) LastGuess() (guess int, ok bool) {
	if len(g.History) == 0 {
		return 0, false
	}
	return g.History[len(g.History)-1], true
}

// GuessesSoFar is the number of guesses made in this round.
func (g *Game) GuessesSoFar() int {
	return len(g.History)
}

// Range returns the inclusive bounds the secret was chosen from.
func (g *Game) Range() (min, max int) {
	return g.Min, g.Max
}
