VERSION?="0.0.1"
GOFMT_FILES?=$$(find . -not -path "./vendor/*" -type f -name '*.go')


default: bin

fmt:
	@gofmt -w $(GOFMT_FILES)

bin:
	@go install ./...

test:
	@go test ./...

protoc:
	@protoc -I pkg/proto/ pkg/proto/service.proto --go_out=plugins=grpc:pkg/proto

.NOTPARALLEL:

.PHONY: bin fmt test protoc