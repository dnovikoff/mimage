.PHONY: all
all: test gobin/mimage

gobin:
	mkdir $@

gobin/mimage: gobin
	go build -o $@ ./cmd/mimage

.PHONY: test
test:
	go test ./pkg/...

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
