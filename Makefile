# output binary file name
output_name = rabbit-topic

# build the source to native OS and platform
.PHONY: build
build:
	go build -ldflags '-extldflags "-static"' -o $(output_name) *.go

# generate proto file to go
output_dir = msg
.PHONY: generate
generate:
	protoc -I . -I ${GOPATH}/src/github.com/googleapis/googleapis -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway --experimental_allow_proto3_optional \
		--go_out ${output_dir} --go_opt paths=source_relative \
		--go-grpc_out ${output_dir} --go-grpc_opt paths=source_relative message.proto
