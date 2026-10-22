set -e
set -o pipefail

#------

rm -rf cmd/mcp-victoriametrics/resources/vmsite

git clone --no-checkout --depth=1 https://github.com/VictoriaMetrics/vmsite.git cmd/mcp-victoriametrics/resources/vmsite
cd cmd/mcp-victoriametrics/resources/vmsite

git sparse-checkout init --cone
git sparse-checkout set content/blog
git checkout master
rm -f ./content/_index.md
find . -mindepth 1 -maxdepth 1 ! -name content -exec rm -rf {} +
find content -type f ! -name '*.md' -delete

cd -

#------
