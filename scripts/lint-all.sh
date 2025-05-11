set -e
set -o pipefail

# check source code by linter
gofmt -l -w -s ./cmd
go vet ./cmd/...
which golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.64.7
golangci-lint run
