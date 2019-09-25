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

# Default GOVERSION
[[ ! "$GOVERSION" ]] && GOVERSION=1.11
REPO=pfcd

testrepo () {
  GO=go
  GO111MODULE=on
  if [[ $GOVERSION == 1.10 ]]; then
    GO=vgo
  fi

  $GO version
  $GO clean -testcache
  $GO build -v ./...
  $GO test -v ./...

  echo "------------------------------------------"
  echo "Tests completed successfully!"
}

DOCKER=
[[ "$1" == "docker" || "$1" == "podman" ]] && DOCKER=$1
if [ ! "$DOCKER" ]; then
    testrepo
    exit
fi

# use Travis cache with docker
DOCKER_IMAGE_TAG=golang-builder-$GOVERSION
mkdir -p ~/.cache
if [ -f ~/.cache/$DOCKER_IMAGE_TAG.tar ]; then
  # load via cache
  $DOCKER load -i ~/.cache/$DOCKER_IMAGE_TAG.tar
else
  # pull and save image to cache
  $DOCKER pull picfight/$DOCKER_IMAGE_TAG
  $DOCKER save picfight/$DOCKER_IMAGE_TAG > ~/.cache/$DOCKER_IMAGE_TAG.tar
fi

$DOCKER run --rm -it -v $(pwd):/src:Z picfight/$DOCKER_IMAGE_TAG /bin/bash -c "\
  rsync -ra --filter=':- .gitignore'  \
  /src/ /go/src/github.com/picfight/$REPO/ && \
  dir && \
  env GOVERSION=$GOVERSION GO111MODULE=on bash run_tests.sh"
