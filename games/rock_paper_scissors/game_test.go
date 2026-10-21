package main

import "testing"

func TestParseMoveValid(t *testing.T) {
	m, err := ParseMove("rock")
	if err != nil || m != Rock {
		t.Fatalf("expected Rock, got %v, err=%v", m, err)
	}
}

func TestParseMoveInvalid(t *testing.T) {
	_, err := ParseMove("lizard")
	if err == nil {
		t.Fatal("expected error for invalid move")
	}
}

func TestBeats(t *testing.T) {
	cases := map[Move]Move{Rock: Scissors, Scissors: Paper, Paper: Rock}
	for move, want := range cases {
		if got := Beats(move); got != want {
			t.Errorf("Beats(%s) = %s, want %s", move, got, want)
		}
	}
}

func TestPlayOutcomes(t *testing.T) {
	cases := []struct {
		player, opponent Move
		want             Outcome
	}{
		{Rock, Scissors, Win},
		{Scissors, Paper, Win},
		{Paper, Rock, Win},
		{Rock, Paper, Lose},
		{Paper, Scissors, Lose},
		{Scissors, Rock, Lose},
		{Rock, Rock, Draw},
		{Paper, Paper, Draw},
		{Scissors, Scissors, Draw},
	}
	for _, c := range cases {
		got := Play(c.player, c.opponent)
		if got != c.want {
			t.Errorf("Play(%s, %s) = %s, want %s", c.player, c.opponent, got, c.want)
		}
	}
}
