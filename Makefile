.PHONY: all
all: test binaries

gobin:
	mkdir $@

gobin/mimage: gobin
	go build -mod vendor  -o $@ ./cmd/mimage

.PHONY: binaries
binaries: gobin/mimage

.PHONY: test
test:
	go test -mod vendor ./...

GO_EXCLUDE := /vendor/
GO_FILES_CMD := find . -name '*.go' | grep -v -E '$(GO_EXCLUDE)'

.PHONY: format
format: $(goimports); $(info $(H) formatting files with goimports...)
	goimports -w -local bitbucket.org/dnovikoff/go-mahjong/mimage $$($(GO_FILES_CMD))
	gofmt -w -s $$($(GO_FILES_CMD))

docker:
	mkdir $@

docker/image: docker Dockerfile
	docker build . -t tempai/mimage > $@

docker/push: docker docker/image
	docker push tempai/mimage > $@
