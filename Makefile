BUILD_DATE := $(shell date +'%d.%m.%Y %H:%M:%S')
BUILD_COMMIT := $(shell git rev-parse --short HEAD)

.PHONY: build
build: build_bot

.PHONY: build_bot
build_bot:
	mkdir -p ./bin
	cd cmd/instadownloadbot && \
	 go build \
	 -ldflags "-X 'main.buildDate=$(BUILD_DATE)' -X main.buildCommit=$(BUILD_COMMIT)" \
	 -o ../../bin/instadownloadbot .
