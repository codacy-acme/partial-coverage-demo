"""Interactive terminal front-end for Tic-Tac-Toe."""
from games.tic_tac_toe.game import InvalidMoveError, TicTacToe


def main():
    game = TicTacToe()
    print("Tic-Tac-Toe! Enter a position 0-8:")
    print(game.board.render())

    while True:
        try:
            raw = input(f"Player {game.current_player}, your move: ")
            position = int(raw)
            game.play(position)
        except (ValueError, InvalidMoveError) as exc:
            print(f"Invalid move: {exc}")
            continue

        print(game.board.render())
        outcome = game.result()
        if outcome:
            print(outcome)
            break


if __name__ == "__main__":
    main()
