# TODO: Build using https://goreleaser.com/
build:
	CGO_ENABLED=0 go build -v

build-linux:
	CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -v

release-linux: build-linux
	tar cvzf urlcrawl-linux-amd64.tar.gz urlcrawl

# TODO: Lint using golangci-lint


