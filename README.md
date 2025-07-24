# MessageApp – gRPC Chat Server in Go

This project implements a simple real-time chat system using Go and gRPC with bidirectional streaming.

## Features

- 🧵 Real-time chat using gRPC bidirectional streams
- 👥 Multi-client support
- 🔧 Cleanly structured with Protobuf
- 🔐 Thread-safe in-memory client management

---

## 📦 Project Structure

MessageApp/
├── client/ # Sample chat client
├── proto/ # Protobuf definition and generated code
├── server/ # ChatService gRPC server implementation
├── main.go # Starts the gRPC server
├── go.mod # Go module
├── README.md # This file
└── Makefile # Build automation (see below)

---

## 🛠️ Requirements

- Go 1.20+
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

Install Go plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
