set -e
set -o pipefail

go build -o ./bin/mcp-victoriametrics ./cmd/mcp-victoriametrics/main.go
