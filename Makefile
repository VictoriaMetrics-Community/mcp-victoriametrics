download-docs:
	bash ./scripts/update-docs.sh

download-blog:
	bash ./scripts/update-blog.sh

update-docs: download-docs

update-blog: download-blog

update-resources: update-docs update-blog

test:
	bash ./scripts/test-all.sh

check:
	bash ./scripts/check-all.sh

lint:
	bash ./scripts/lint-all.sh

build:
	bash ./scripts/build-binaries.sh

all: test check lint build
