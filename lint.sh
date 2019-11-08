#!/usr/bin/env bash

set -e

GOLANGCI_LINT_VERSION=1.21.0

LINTER=${GOPATH}/bin/golangci-lint

if [ ! -f "$LINTER" ]; then
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v$GOLANGCI_LINT_VERSION
fi


${LINTER} run -D structcheck \
    -E bodyclose \
    -E depguard \
    -E dupl \
    -E goconst \
    -E gocritic \
    -E gocyclo \
    -E gofmt \
    -E goimports \
    -E gosec \
    -E interfacer \
    -E maligned \
    -E prealloc \
    -E scopelint \
    -E stylecheck \
    -E unconvert \
    -E unparam \
    -e "field.* is unused" \
    ./...
