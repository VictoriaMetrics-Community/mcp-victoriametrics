set -e
set -o pipefail

# check licenses
which wwhrd || go install github.com/frapposelli/wwhrd@latest
wwhrd check -f .wwhrd.yml

# check for vulnerabilities
which govulncheck || go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
