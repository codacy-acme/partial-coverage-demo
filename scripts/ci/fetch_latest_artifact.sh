#!/usr/bin/env bash
# Fetch the most recent non-expired GitHub Actions artifact with the given
# name, from anywhere in this repo's workflow history, and unzip it into
# out_dir. This is how an untouched module's coverage is "retrieved" rather
# than regenerated: this run's own workflow already uploaded a fresh
# artifact for any module it just tested, so it naturally sorts newest and
# gets picked here too — the same lookup covers both cases.
#
# Usage: fetch_latest_artifact.sh <artifact-name> <out-dir>
set -euo pipefail

artifact_name=$1
out_dir=$2
mkdir -p "$out_dir"

if [ -z "${GITHUB_REPOSITORY:-}" ]; then
  echo "GITHUB_REPOSITORY must be set" >&2
  exit 1
fi

response=$(gh api "repos/${GITHUB_REPOSITORY}/actions/artifacts?per_page=100")

candidates=$(echo "$response" | jq -r --arg name "$artifact_name" '
  [.artifacts[] | select(.name == $name and .expired == false)] | length
')

selected=$(echo "$response" | jq -r --arg name "$artifact_name" '
  .artifacts
  | map(select(.name == $name and .expired == false))
  | sort_by(.created_at)
  | reverse
  | .[0]
  | if . == null then "" else "\(.id)\t\(.created_at)" end
')

if [ -z "$selected" ]; then
  echo "No artifact named '$artifact_name' found yet (expected before that module's first CI run) — skipping." >&2
  exit 0
fi

artifact_id=${selected%%$'\t'*}
created_at=${selected#*$'\t'}
echo "Found $candidates candidate(s) for '$artifact_name'; using the newest (id $artifact_id, created $created_at)."
zip_path="$out_dir/$artifact_name.zip"
gh api "repos/${GITHUB_REPOSITORY}/actions/artifacts/${artifact_id}/zip" > "$zip_path"
unzip -o -q "$zip_path" -d "$out_dir"
rm "$zip_path"
