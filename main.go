package main

import (
	"log"
	"net"

	pb "github.com/pgibb96/MessageApp/proto"
	"github.com/pgibb96/MessageApp/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chatServer := server.NewServer()

	pb.RegisterChatServiceServer(grpcServer, chatServer)

	log.Println("🚀 Chat server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
