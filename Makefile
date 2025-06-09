download-docs:
	bash ./scripts/update-docs.sh

update-docs: download-docs

test:
	bash ./scripts/test-all.sh

check:
	bash ./scripts/check-all.sh

lint:
	bash ./scripts/lint-all.sh

build:
	bash ./scripts/build-binaries.sh

all: test check lint build
