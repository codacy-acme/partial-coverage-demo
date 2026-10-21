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

func TestOpposite(t *testing.T) {
	cases := map[Outcome]Outcome{Win: Lose, Lose: Win, Draw: Draw}
	for in, want := range cases {
		if got := Opposite(in); got != want {
			t.Errorf("Opposite(%s) = %s, want %s", in, got, want)
		}
	}
}

func TestRandomMoveIsAlwaysValid(t *testing.T) {
	for i := 0; i < 50; i++ {
		m := RandomMove()
		if !validMoves[m] {
			t.Fatalf("RandomMove() returned invalid move %s", m)
		}
	}
}

func TestIsValidMove(t *testing.T) {
	for _, m := range AllMoves {
		if !IsValidMove(string(m)) {
			t.Errorf("IsValidMove(%q) = false, want true", m)
		}
	}
	if IsValidMove("lizard") {
		t.Error("IsValidMove(\"lizard\") = true, want false")
	}
}

func TestAllMoves(t *testing.T) {
	if len(AllMoves) != 3 {
		t.Fatalf("expected 3 moves, got %d", len(AllMoves))
	}
	for _, m := range AllMoves {
		if !validMoves[m] {
			t.Errorf("AllMoves contains invalid move %s", m)
		}
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
