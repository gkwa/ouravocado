#!/bin/bash

for i in {1..50}; do
  depth=$((RANDOM % 10 + 1))
  dir=testdata
  for j in $(seq 1 $depth); do
    subdir="dir$((RANDOM % 5 + 1))"
    dir="$dir/$subdir"
  done

  mkdir -p "$dir"
  file="$dir/file$i.md"
  size=$((RANDOM % 10000000000 + 1))
  head -c "$size" /dev/urandom | tr -dc 'a-zA-Z0-9 ' | fold -w 80 > "$file"
done
