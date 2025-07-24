package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/pgibb96/MessageApp/proto"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run client.go <your-name>")
	}
	username := os.Args[1]

	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	// Start ChatStream
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("Error starting chat stream: %v", err)
	}

	// Send initial join message
	err = stream.Send(&pb.MessageRequest{
		Sender:  username,
		Message: fmt.Sprintf("%s has joined the chat.", username),
	})
	if err != nil {
		log.Fatalf("Error sending join message: %v", err)
	}

	// Goroutine to receive messages
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			ts := time.Unix(msg.Timestamp, 0).Format("15:04:05")
			fmt.Printf("[%s] %s: %s\n", ts, msg.Sender, msg.Message)
		}
	}()

	// Read user input and send messages
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages (Ctrl+C to exit):")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		err = stream.Send(&pb.MessageRequest{
			Sender:  username,
			Message: text,
		})
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	}

	if scanner.Err() != nil {
		log.Printf("Scanner error: %v", scanner.Err())
	}
}
