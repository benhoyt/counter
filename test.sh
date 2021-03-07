#!/usr/bin/env bash

set -e

cat kjvbible.txt kjvbible.txt kjvbible.txt kjvbible.txt kjvbible.txt kjvbible.txt kjvbible.txt kjvbible.txt kjvbible.txt kjvbible.txt >kjvbible_x10.txt

go build ./cmd/bench

echo Testing map output against baseline
./bench map <kjvbible_x10.txt | python3 normalize.py >output.txt
git diff --exit-code output.txt

echo Testing counter output against baseline
./bench counter <kjvbible_x10.txt | python3 normalize.py >output.txt
git diff --exit-code output.txt

echo Benchmarking map version
for i in {1..5}
do
  time ./bench map <kjvbible_x10.txt >/dev/null
done

echo Benchmarking counter version
for i in {1..5}
do
  time ./bench counter <kjvbible_x10.txt >/dev/null
done
