package main

import "testing"

func TestNewDeckHas52Cards(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 52 {
		t.Fatalf("expected 52 cards, got %d", len(deck))
	}
}

func TestHandValueSimple(t *testing.T) {
	hand := []Card{{Rank: "10", Suit: "♠"}, {Rank: "7", Suit: "♥"}}
	if v := HandValue(hand); v != 17 {
		t.Errorf("expected 17, got %d", v)
	}
}

func TestHandValueSoftAce(t *testing.T) {
	hand := []Card{{Rank: "A", Suit: "♠"}, {Rank: "9", Suit: "♥"}}
	if v := HandValue(hand); v != 20 {
		t.Errorf("expected soft ace to score 20, got %d", v)
	}
}

func TestHandValueBustedAceCountsAsOne(t *testing.T) {
	hand := []Card{{Rank: "A", Suit: "♠"}, {Rank: "9", Suit: "♥"}, {Rank: "5", Suit: "♦"}}
	if v := HandValue(hand); v != 15 {
		t.Errorf("expected ace to count as 1 once it would bust, got %d", v)
	}
}

func TestIsBust(t *testing.T) {
	bust := []Card{{Rank: "K", Suit: "♠"}, {Rank: "Q", Suit: "♥"}, {Rank: "5", Suit: "♦"}}
	if !IsBust(bust) {
		t.Error("expected hand to be bust")
	}
	ok := []Card{{Rank: "K", Suit: "♠"}, {Rank: "5", Suit: "♥"}}
	if IsBust(ok) {
		t.Error("expected hand not to be bust")
	}
}

func TestDetermineWinner(t *testing.T) {
	playerBust := []Card{{Rank: "K", Suit: "♠"}, {Rank: "Q", Suit: "♥"}, {Rank: "5", Suit: "♦"}}
	dealerOk := []Card{{Rank: "10", Suit: "♠"}, {Rank: "5", Suit: "♥"}}
	if r := DetermineWinner(playerBust, dealerOk); r != DealerWins {
		t.Errorf("expected DealerWins on player bust, got %v", r)
	}

	playerOk := []Card{{Rank: "10", Suit: "♠"}, {Rank: "9", Suit: "♥"}}
	dealerBust := []Card{{Rank: "K", Suit: "♠"}, {Rank: "Q", Suit: "♥"}, {Rank: "5", Suit: "♦"}}
	if r := DetermineWinner(playerOk, dealerBust); r != PlayerWins {
		t.Errorf("expected PlayerWins on dealer bust, got %v", r)
	}

	push := []Card{{Rank: "10", Suit: "♠"}, {Rank: "9", Suit: "♥"}}
	if r := DetermineWinner(playerOk, push); r != PushResult {
		t.Errorf("expected push on equal scores, got %v", r)
	}

	higher := []Card{{Rank: "10", Suit: "♠"}, {Rank: "K", Suit: "♥"}}
	if r := DetermineWinner(higher, playerOk); r != PlayerWins {
		t.Errorf("expected PlayerWins on higher score, got %v", r)
	}
}

func TestDealerShouldHit(t *testing.T) {
	low := []Card{{Rank: "10", Suit: "♠"}, {Rank: "5", Suit: "♥"}}
	if !DealerShouldHit(low) {
		t.Error("expected dealer to hit on 15")
	}
	high := []Card{{Rank: "10", Suit: "♠"}, {Rank: "9", Suit: "♥"}}
	if DealerShouldHit(high) {
		t.Error("expected dealer to stand on 19")
	}
}
