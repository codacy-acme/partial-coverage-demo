"""Core Hangman game logic, independent of any I/O."""

DEFAULT_WORDS = ["python", "codacy", "monorepo", "coverage", "artifact"]


class HangmanGame:
    def __init__(self, word, max_misses=6):
        if not word:
            raise ValueError("word must be non-empty")
        self.word = word.lower()
        self.max_misses = max_misses
        self.guessed = set()
        self.misses = 0

    def guess(self, letter):
        letter = letter.lower()
        if letter in self.guessed:
            return False
        self.guessed.add(letter)
        if letter not in self.word:
            self.misses += 1
            return False
        return True

    @property
    def masked_word(self):
        return "".join(ch if ch in self.guessed else "_" for ch in self.word)

    def is_won(self):
        return all(ch in self.guessed for ch in self.word)

    def is_lost(self):
        return self.misses >= self.max_misses

    def is_over(self):
        return self.is_won() or self.is_lost()

    def remaining_attempts(self):
        return self.max_misses - self.misses

    @property
    def incorrect_guesses(self):
        return sorted(letter for letter in self.guessed if letter not in self.word)

    @property
    def guesses_made(self):
        return len(self.guessed)

    def __str__(self):
        return self.masked_word

    @property
    def word_length(self):
        return len(self.word)

    def give_up(self):
        self.misses = self.max_misses
