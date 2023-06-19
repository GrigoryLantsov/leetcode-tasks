#!/bin/bash

# Input array
gain=("$@")

# Calculate altitudes
altitude=0
max_altitude=0

for i in "${gain[@]}"; do
  altitude=$((altitude + i))
  if (( altitude > max_altitude )); then
    max_altitude=$altitude
  fi
done

echo "The highest altitude is $max_altitude"
