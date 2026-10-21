import pytest

from games.hangman.game import HangmanGame


def test_empty_word_raises():
    with pytest.raises(ValueError):
        HangmanGame("")


def test_masked_word_starts_fully_hidden():
    game = HangmanGame("cat")
    assert game.masked_word == "___"


def test_correct_guess_reveals_letter():
    game = HangmanGame("cat")
    assert game.guess("c") is True
    assert game.masked_word == "c__"
    assert game.misses == 0


def test_incorrect_guess_increments_misses():
    game = HangmanGame("cat", max_misses=3)
    assert game.guess("z") is False
    assert game.misses == 1


def test_repeated_guess_is_noop():
    game = HangmanGame("cat")
    game.guess("c")
    assert game.guess("c") is False
    assert game.misses == 0


def test_win_condition():
    game = HangmanGame("cat")
    for letter in "cat":
        game.guess(letter)
    assert game.is_won()
    assert game.is_over()
    assert not game.is_lost()


def test_guesses_made_counts_all_unique_letters():
    game = HangmanGame("cat")
    game.guess("c")
    game.guess("z")
    game.guess("c")
    assert game.guesses_made == 2


def test_incorrect_guesses_tracks_misses_only():
    game = HangmanGame("cat")
    game.guess("c")
    game.guess("z")
    game.guess("y")
    assert game.incorrect_guesses == ["y", "z"]


def test_lose_condition():
    game = HangmanGame("cat", max_misses=2)
    game.guess("x")
    game.guess("y")
    assert game.is_lost()
    assert game.is_over()
    assert not game.is_won()
    assert game.remaining_attempts() == 0
