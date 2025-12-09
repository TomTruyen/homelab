#!/usr/bin/env bash
set -euo pipefail

tmpfile="$(mktemp)"

# Header
cat > "$tmpfile" <<'YAML'
version: 2
updates:
  - package-ecosystem: "docker-compose"
    directories:
YAML

# Find all docker-compose.yml / docker-compose.yaml files anywhere
# Store directories in an array to avoid subshell issues
mapfile -d '' dirs < <(find . -type f \( -name "docker-compose.yml" -o -name "docker-compose.yaml" \) -print0)

# Deduplicate directories
declare -A seen
for file in "${dirs[@]}"; do
    echo "Processing file: $file" >&2

    dir="$(dirname "$file")"
    dir="${dir#./}"  # remove leading ./
    if [[ -z "${seen[$dir]+x}" ]]; then
        echo "      - \"/$dir\"" >> "$tmpfile"
        seen["$dir"]=1
    fi
done

# Append the schedule block
cat >> "$tmpfile" <<'YAML'
    schedule:
      interval: "weekly"
      day: "sunday"
      time: "06:00"
      timezone: "Europe/Brussels"
YAML

# Install if changed
if ! [ -f .github/dependabot.yml ] || ! cmp -s "$tmpfile" .github/dependabot.yml; then
  mv "$tmpfile" .github/dependabot.yml
  echo "✅ Updated .github/dependabot.yml!"
else
  echo "ℹ️ No changes to .github/dependabot.yml."
fi
