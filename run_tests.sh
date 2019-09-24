#!/usr/bin/env bash

GO=go
GO111MODULE=on

  $GO version
  $GO clean -testcache
  $GO build -v ./...
  $GO test -v ./...
