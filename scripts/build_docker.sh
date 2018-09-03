#!/usr/bin/env sh

set -e

env GOOS=linux GOARCH=arm go build github.com/ekrengel/introk8s
docker build .
