package main

import (
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	pb "github.com/pgibb96/MessageApp/proto"
	"github.com/pgibb96/MessageApp/server"
)

func main() {
	// 1. Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// 2. Get JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("❌ Failed to initialize JetStream: %v", err)
	}

	// 3. Ensure the stream exists
	_, err = js.StreamInfo("CHAT")
	if err != nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     "CHAT",
			Subjects: []string{"chat.*"},
			Storage:  nats.FileStorage, // persistent
			Retention: nats.LimitsPolicy,
		})
		if err != nil {
			log.Fatalf("❌ Failed to create JetStream stream: %v", err)
		}
		log.Println("📦 JetStream stream 'CHAT' created")
	} else {
		log.Println("📦 JetStream stream 'CHAT' already exists")
	}

	// 4. Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, server.NewServer(nc, js))

	log.Println("🚀 gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}