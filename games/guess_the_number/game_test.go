package main

import "testing"

func TestNewGameInvalidSecret(t *testing.T) {
	if _, err := NewGame(200, 1, 100, 10); err == nil {
		t.Fatal("expected error for out-of-range secret")
	}
}

func TestGuessHints(t *testing.T) {
	g, err := NewGame(50, 1, 100, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if hint := g.Guess(10); hint != Higher {
		t.Errorf("expected Higher, got %v", hint)
	}
	if hint := g.Guess(90); hint != Lower {
		t.Errorf("expected Lower, got %v", hint)
	}
	if hint := g.Guess(50); hint != Correct {
		t.Errorf("expected Correct, got %v", hint)
	}
	if g.Attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", g.Attempts)
	}
}

func TestLastGuess(t *testing.T) {
	g, _ := NewGame(50, 1, 100, 10)
	if _, ok := g.LastGuess(); ok {
		t.Fatal("expected no last guess before any guess made")
	}
	g.Guess(10)
	g.Guess(30)
	if last, ok := g.LastGuess(); !ok || last != 30 {
		t.Errorf("LastGuess() = (%d, %v), want (30, true)", last, ok)
	}
}

func TestReset(t *testing.T) {
	g, _ := NewGame(50, 1, 100, 10)
	g.Guess(10)
	g.Guess(20)
	g.Reset(75)
	if g.Secret != 75 {
		t.Errorf("expected secret 75 after reset, got %d", g.Secret)
	}
	if g.Attempts != 0 {
		t.Errorf("expected 0 attempts after reset, got %d", g.Attempts)
	}
	if len(g.History) != 0 {
		t.Errorf("expected empty history after reset, got %v", g.History)
	}
}

func TestIsWithinRange(t *testing.T) {
	g, _ := NewGame(50, 1, 100, 10)
	if !g.IsWithinRange(1) || !g.IsWithinRange(100) || !g.IsWithinRange(50) {
		t.Error("expected bounds and midpoint to be within range")
	}
	if g.IsWithinRange(0) || g.IsWithinRange(101) {
		t.Error("expected values outside [1, 100] to not be within range")
	}
}

func TestRange(t *testing.T) {
	g, _ := NewGame(50, 1, 100, 10)
	if min, max := g.Range(); min != 1 || max != 100 {
		t.Errorf("Range() = (%d, %d), want (1, 100)", min, max)
	}
}

func TestGuessesSoFar(t *testing.T) {
	g, _ := NewGame(50, 1, 100, 10)
	if g.GuessesSoFar() != 0 {
		t.Fatalf("expected 0 guesses before any guess made, got %d", g.GuessesSoFar())
	}
	g.Guess(10)
	g.Guess(20)
	if g.GuessesSoFar() != 2 {
		t.Errorf("expected 2 guesses, got %d", g.GuessesSoFar())
	}
}

func TestHistoryTracksGuesses(t *testing.T) {
	g, _ := NewGame(50, 1, 100, 10)
	g.Guess(10)
	g.Guess(90)
	g.Guess(50)
	want := []int{10, 90, 50}
	if len(g.History) != len(want) {
		t.Fatalf("expected history %v, got %v", want, g.History)
	}
	for i, v := range want {
		if g.History[i] != v {
			t.Errorf("history[%d] = %d, want %d", i, g.History[i], v)
		}
	}
}

func TestIsOverAndAttemptsRemaining(t *testing.T) {
	g, _ := NewGame(50, 1, 100, 2)
	if g.IsOver() {
		t.Fatal("game should not be over yet")
	}
	g.Guess(1)
	if g.AttemptsRemaining() != 1 {
		t.Errorf("expected 1 attempt remaining, got %d", g.AttemptsRemaining())
	}
	g.Guess(2)
	if !g.IsOver() {
		t.Fatal("expected game to be over after max attempts")
	}
}
