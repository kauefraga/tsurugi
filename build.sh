#! /usr/bin/bash

echo "building tsurugi for the following platforms: linux and darwin"

if [ ! -d "bin" ]; then
  mkdir bin
fi

CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -o bin/tsurugi cmd/main.go
CGO_ENABLED=0 GOOS=darwin go build -ldflags='-w -s' -o bin/tsurugi-darwin cmd/main.go

echo "finished, look in the bin/ folder"
