#!/usr/bin/env python3
"""Launcher for the Python games in this monorepo.

Usage: python play.py <game>
"""
import importlib
import sys

PYTHON_GAMES = {
    "tic-tac-toe": "games.tic_tac_toe.cli",
    "hangman": "games.hangman.cli",
}


def main():
    if len(sys.argv) != 2 or sys.argv[1] not in PYTHON_GAMES:
        print("Usage: python play.py <game>")
        print(f"Available games: {', '.join(sorted(PYTHON_GAMES))}")
        sys.exit(1)

    module = importlib.import_module(PYTHON_GAMES[sys.argv[1]])
    module.main()


if __name__ == "__main__":
    main()
