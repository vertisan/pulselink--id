#!/bin/sh
set -e

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin v1.53.3

go mod vendor
go mod verify

golangci-lint --version
golangci-lint run
