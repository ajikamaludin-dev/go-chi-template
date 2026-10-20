#!/usr/bin/env bash
set -euo pipefail

if [ -z "${1:-}" ]; then
  echo "Usage: $0 <module-path>  (e.g. github.com/acme/my-service)"
  exit 1
fi

NEW="$1"
OLD="github.com/aji/chi-go-template"

# update go.mod module line
go mod edit -module "$NEW"

# replace import paths across all .go files
grep -rl "$OLD" --include='*.go' . | while read -r f; do
  sed -i "s|$OLD|$NEW|g" "$f"
done

go mod tidy
echo "Module renamed to $NEW"