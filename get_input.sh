#!/usr/bin/env bash

set -e

export "$(xargs < .env)"

export YEAR=$1
export DAY=$2

function get_input() {
  input_folder="inputs/$YEAR"
  input_file="$input_folder/$DAY"

  if [ ! -f "$input_file" ]; then
    mkdir -p "$input_folder"
    source ./.env
    >&2 echo "downloading input from https://adventofcode.com/$YEAR/day/$((DAY))/input"
    curl -s "https://adventofcode.com/$YEAR/day/$((DAY))/input" -H "Cookie: session=$SESSION" > "$input_file"
  fi
}

get_input