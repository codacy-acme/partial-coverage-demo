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
  tic_tac_toe/   (Python)
  hangman/       (Python)
  ...more games added in later PRs (Go)
```

Each game under `games/<name>/` is a self-contained module with its own
game logic, a terminal front-end, and its own tests — so each module's
coverage can be generated and uploaded independently.

## Playing the games

```
pip install -r requirements-dev.txt
python play.py tic-tac-toe
python play.py hangman
```

## Running tests / coverage for a single module

```
coverage run --source=games/tic_tac_toe -m pytest games/tic_tac_toe
coverage xml -o coverage-tic-tac-toe.xml
```

CI wiring to generate these per-module reports automatically, cache them
as artifacts, and merge + upload them to Codacy lands in follow-up PRs.
