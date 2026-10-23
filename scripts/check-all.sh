set -e
set -o pipefail

# check licenses
which wwhrd || go install github.com/frapposelli/wwhrd@latest
wwhrd check -f .wwhrd.yml

# check for vulnerabilities
which govulncheck || go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...

# check that generated sections of README.md and web UI client configs are up to date
go run ./scripts/readmegen
git diff --exit-code README.md web/src/clients.gen.json || (echo "README.md or web/src/clients.gen.json is outdated, run 'make update-readme' and commit the result" && exit 1)
