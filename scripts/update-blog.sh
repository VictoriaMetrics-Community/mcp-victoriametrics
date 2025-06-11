set -e
set -o pipefail

#------

rm -rf cmd/mcp-victoriametrics/resources/vmsite

git clone --no-checkout --depth=1 https://github.com/VictoriaMetrics/vmsite.git cmd/mcp-victoriametrics/resources/vmsite
cd cmd/mcp-victoriametrics/resources/vmsite

git sparse-checkout init --cone
git sparse-checkout set content/blog
git checkout master
rm -rf ./.git
rm -f ./content/_index.md ./Dockerfile ./Makefile ./*.md ./*.json ./*.lock ./.gitignore

cd -

#------
