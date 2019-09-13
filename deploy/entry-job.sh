#!/bin/bash

spacemarket &
spacemarket_pid=$!
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start spacemarket: $status"
  exit $status
fi

trap "func_trap" "TERM"

func_trap() {
    kill $spacemarket_pid
}

while :
do
  sleep 1
done
