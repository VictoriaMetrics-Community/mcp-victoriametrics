set -e
set -o pipefail

#------

rm -rf cmd/mcp-victoriametrics/resources/vm

git clone --no-checkout --depth=1 https://github.com/VictoriaMetrics/vmdocs.git cmd/mcp-victoriametrics/resources/vm
cd cmd/mcp-victoriametrics/resources/vm

git sparse-checkout init --cone
git sparse-checkout set content
git checkout main
rm -rf ./.git
rm -f ./docs/Makefile ./Makefile ./LICENSE ./*.md ./*.mod ./*.sum ./*.zip ./.golangci.yml ./.wwhrd.yml ./.gitignore ./.dockerignore ./codecov.yml ./Dockerfile ./*.sh ./*.js ./*.json ./*.lock

cd -

#------
