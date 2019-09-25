#!/usr/bin/env bash

# usage:
# ./run_tests.sh                         # local, go 1.11
# GOVERSION=1.10 ./run_tests.sh          # local, go 1.10 (vgo)
# ./run_tests.sh docker                  # docker, go 1.11
# GOVERSION=1.10 ./run_tests.sh docker   # docker, go 1.10 (vgo)
# ./run_tests.sh podman                  # podman, go 1.11
# GOVERSION=1.10 ./run_tests.sh podman   # podman, go 1.10 (vgo)

set -ex

# The script does automatic checking on a Go package and its sub-packages,
# including:
# 1. gofmt         (http://golang.org/cmd/gofmt/)
# 2. gosimple      (https://github.com/dominikh/go-simple)
# 3. unconvert     (https://github.com/mdempsky/unconvert)
# 4. ineffassign   (https://github.com/gordonklaus/ineffassign)
# 5. race detector (http://blog.golang.org/race-detector)

# gometalinter (github.com/alecthomas/gometalinter) is used to run each each
# static checker.

# To run on docker on windows, symlink /mnt/c to /c and then execute the script
# from the repo path under /c.  See:
# https://github.com/Microsoft/BashOnWindows/issues/1854
# for more details.

  GO111MODULE=on

  go version
  go env
  go clean -testcache
  go build -v ./...
  go test -v ./...