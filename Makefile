PROTO_SRC=proto/chat.proto
PROTO_DIR=proto
PROTO_PATH=proto

proto:
	protoc \
		--proto_path=$(PROTO_PATH) \
		--go_out=$(PROTO_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_SRC)

clean:
	rm -f $(PROTO_DIR)/*.pb.go

format:
	go fmt ./...

run-server:
	go run main.go

test:
	go test ./...