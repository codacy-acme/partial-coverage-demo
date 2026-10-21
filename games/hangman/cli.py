"""Interactive terminal front-end for Hangman."""
import random

from games.hangman.game import DEFAULT_WORDS, HangmanGame


def main():
    word = random.choice(DEFAULT_WORDS)
    game = HangmanGame(word)
    print("Hangman! Guess the word one letter at a time.")

    while not game.is_over():
        print(f"\nWord: {game.masked_word}")
        print(f"Attempts left: {game.remaining_attempts()}")
        letter = input("Guess a letter: ").strip()
        if len(letter) != 1 or not letter.isalpha():
            print("Please enter a single letter.")
            continue
        hit = game.guess(letter)
        print("Correct!" if hit else "Nope.")

    if game.is_won():
        print(f"\nYou win! The word was '{game.word}'.")
    else:
        print(f"\nYou lose! The word was '{game.word}'.")


if __name__ == "__main__":
    main()
