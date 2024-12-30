### Protobuf Tools
BUF_VERSION := v1.48.0
PROTOC_GEN_GO_VERSION := v1.36.1
PROTOC_GEN_GO_GRPC_VERSION := v1.5.1

## Variables
mk_path  := $(abspath $(lastword $(MAKEFILE_LIST)))
root_dir   := $(dir $(mk_path))
proto_dir := $(root_dir)/api/proto
tool_bin := $(root_dir)/bin

## Targets
BUF := $(tool_bin)/buf
$(BUF):
	@echo "Install proto plugins to $(tool_bin)"
	@mkdir -p $(tool_bin)
	@rm -f $(tool_bin)/protoc-gen-go
	@rm -f $(tool_bin)/protoc-gen-go-grpc
	@rm -f $(tool_bin)/buf
	@GOBIN=$(tool_bin) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@GOBIN=$(tool_bin) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	@GOBIN=$(tool_bin) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

.PHONY: generate
generate: $(BUF)
	@PATH=$(tool_bin):$(proto_dir) $(BUF) generate
