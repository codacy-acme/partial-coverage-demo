# partial-coverage-demo

A monorepo of small classic games, used to demonstrate **partial/per-module
coverage reporting to Codacy** in a monolith-style codebase.

## The idea

This repo is a stand-in for a monolith with many independently-owned
modules. The problem it demonstrates:

- When you touch module `x`, CI only re-runs tests and regenerates coverage
  for `x`.
- Coverage for every other module (`a`..`y`) is *not* regenerated — it's
  fetched from the most recent GitHub Actions artifact produced for that
  module (e.g. from `main`).
- All per-module coverage reports are merged and uploaded to Codacy as a
  single, complete picture of the repo, even though only one module was
  actually re-tested in this run.

## Layout

```
games/
  tic_tac_toe/           (Python)
  hangman/                (Python)
  rock_paper_scissors/    (Go)
  guess_the_number/       (Go)
  blackjack/              (Go)
```

Each game under `games/<name>/` is a self-contained module (own game
logic, terminal front-end, and tests) — so each module's coverage can be
generated and uploaded independently. Go games are each their own Go
module (`go.mod`) for the same reason.

## Playing the games

```
pip install -r requirements-dev.txt
python play.py tic-tac-toe
python play.py hangman

cd games/rock_paper_scissors && go run .
cd games/guess_the_number && go run .
cd games/blackjack && go run .
```

## Running tests / coverage for a single module

Python:

```
coverage run --source=games/tic_tac_toe -m pytest games/tic_tac_toe
coverage xml -o coverage-tic-tac-toe.xml
```

Go:

```
cd games/rock_paper_scissors
go test ./... -coverprofile=coverage.out
```

## CI

`.github/workflows/coverage.yml` does this automatically on every push
and PR:

1. Detects which `games/<name>/` module(s) changed.
2. Tests and uploads a `coverage-<module>` artifact only for those.
3. Fetches the most recent `coverage-<module>` artifact for every
   module (this run's, if it was just regenerated above; otherwise the
   latest one on record) and uploads all of them to Codacy as partial
   reports for the commit, then finalizes.

To force a full run across all 5 modules (e.g. to seed or refresh the
baseline), trigger it manually:

```
gh workflow run coverage.yml --ref main
```
