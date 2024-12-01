#!/bin/bash

if [ $# -ne 1 ]; then
  echo "Usage: ./run_solution.sh <day-folder>"
  exit 1
fi

DAY_FOLDER=$1

if [ ! -d "$DAY_FOLDER" ]; then
  echo "Folder $DAY_FOLDER does not exist!"
  exit 1
fi

cd "$DAY_FOLDER"

IMAGE_TAG="aoc-${DAY_FOLDER}"

docker build -q -t "$IMAGE_TAG" .

docker run --rm "$IMAGE_TAG"

docker rmi "$IMAGE_TAG"
