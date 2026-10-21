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
