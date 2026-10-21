import pytest

from games.tic_tac_toe.game import Board, InvalidMoveError, TicTacToe


def test_board_starts_empty():
    board = Board()
    assert board.cells == [" "] * 9
    assert board.winner() is None
    assert not board.is_full()


def test_place_and_render():
    board = Board()
    board.place(0, "X")
    assert board.cells[0] == "X"
    assert "X" in board.render()


def test_place_out_of_range():
    board = Board()
    with pytest.raises(InvalidMoveError):
        board.place(9, "X")


def test_place_on_taken_cell():
    board = Board()
    board.place(0, "X")
    with pytest.raises(InvalidMoveError):
        board.place(0, "O")


def test_row_win():
    board = Board()
    for pos in (0, 1, 2):
        board.place(pos, "X")
    assert board.winner() == "X"
    assert board.is_game_over()


def test_draw():
    board = Board()
    # X O X / X O O / O X X -> no winner, full board
    marks = ["X", "O", "X", "X", "O", "O", "O", "X", "X"]
    for pos, mark in enumerate(marks):
        board.place(pos, mark)
    assert board.winner() is None
    assert board.is_full()


def test_full_game_win_flow():
    game = TicTacToe()
    assert game.current_player == "X"
    game.play(0)  # X
    game.play(3)  # O
    game.play(1)  # X
    game.play(4)  # O
    game.play(2)  # X wins top row
    assert game.result() == "X wins!"


def test_play_after_game_over_raises():
    game = TicTacToe()
    game.play(0)
    game.play(3)
    game.play(1)
    game.play(4)
    game.play(2)  # X wins
    with pytest.raises(InvalidMoveError):
        game.play(5)


def test_draw_game_result():
    game = TicTacToe()
    for pos in (0, 1, 2, 5, 3, 4, 7, 6, 8):
        game.play(pos)
    assert game.result() == "Draw!"
