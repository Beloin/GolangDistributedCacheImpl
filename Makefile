project := proto build

build:
	go build cmd/main.go

proto:
	protoc -I=internal/network/proto --go_out=. --go_opt=module=beloin.com/distributed-cache \
		--go-grpc_out=. --go-grpc_opt=module=beloin.com/distributed-cache \
		internal/network/proto/*.proto

clean: clean_pb

pb_files = $(shell find ./ -name "*.pb.*" -type f)
clean_pb:
	rm -f $(pb_files)

all: $(project)

test: all
	go test ./...
