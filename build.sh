#!/bin/sh

cd "$(dirname "$0")" || exit

COMMIT_HASH=$(git rev-parse --short HEAD)
TAG=$(git describe --tags --abbrev=0)

go build -ldflags "-X main.CommitHash=$COMMIT_HASH -X main.Tag=$TAG" -o mp main.go
