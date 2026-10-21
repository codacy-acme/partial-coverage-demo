"""Core Tic-Tac-Toe game logic, independent of any I/O."""

WINNING_LINES = [
    (0, 1, 2), (3, 4, 5), (6, 7, 8),  # rows
    (0, 3, 6), (1, 4, 7), (2, 5, 8),  # columns
    (0, 4, 8), (2, 4, 6),             # diagonals
]


class InvalidMoveError(Exception):
    pass


class Board:
    def __init__(self):
        self.cells = [" "] * 9

    def reset(self):
        self.cells = [" "] * 9

    def place(self, position, mark):
        if not 0 <= position <= 8:
            raise InvalidMoveError(f"Position {position} is out of range (0-8)")
        if self.cells[position] != " ":
            raise InvalidMoveError(f"Position {position} is already taken")
        self.cells[position] = mark

    def winner(self):
        for a, b, c in WINNING_LINES:
            if self.cells[a] != " " and self.cells[a] == self.cells[b] == self.cells[c]:
                return self.cells[a]
        return None

    def is_full(self):
        return " " not in self.cells

    def available_moves(self):
        return [i for i, cell in enumerate(self.cells) if cell == " "]

    def corners(self):
        return [self.cells[i] for i in (0, 2, 6, 8)]

    def center(self):
        return self.cells[4]

    def is_game_over(self):
        return self.winner() is not None or self.is_full()

    def render(self):
        rows = [self.cells[i:i + 3] for i in range(0, 9, 3)]
        return "\n------\n".join("|".join(row) for row in rows)

    def __str__(self):
        return self.render()


class TicTacToe:
    PLAYERS = ("X", "O")

    def __init__(self):
        self.board = Board()
        self.turn_index = 0

    @property
    def current_player(self):
        return self.PLAYERS[self.turn_index]

    @property
    def moves_played(self):
        return 9 - len(self.board.available_moves())

    def play(self, position):
        if self.board.is_game_over():
            raise InvalidMoveError("Game is already over")
        self.board.place(position, self.current_player)
        self.turn_index = 1 - self.turn_index

    def result(self):
        winner = self.board.winner()
        if winner:
            return f"{winner} wins!"
        if self.board.is_full():
            return "Draw!"
        return None
