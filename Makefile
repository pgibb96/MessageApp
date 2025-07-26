PROTO_SRC=proto/chat.proto
PROTO_DIR=proto
PROTO_PATH=proto

proto:
	buf generate

clean:
	rm -f $(PROTO_DIR)/*.pb.go

format:
	go fmt ./...

run-server:
	go run main.go

test:
	go test ./...