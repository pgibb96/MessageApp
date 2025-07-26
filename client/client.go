package main
import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	pb "github.com/pgibb96/MessageApp/proto"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run client.go <your-name>")
	}
	username := os.Args[1]

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("Failed to create chat stream: %v", err)
	}

	// Track last joined channel for convenience
	var lastChannel string

	// Start goroutine to receive messages
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Stream receive error: %v", err)
			}
	
			// Clear current line and move cursor to beginning
			fmt.Print("\r\033[K")
	
			// Color and format message
			ts := time.Unix(msg.Timestamp, 0).Format("15:04:05")
			header := color.New(color.FgCyan).Sprintf("[%s] #%s %s:", ts, msg.Channel, msg.Sender)
			fmt.Printf("%s %s\n", header, msg.Message)
	
			// Reprint prompt
			fmt.Print("> ")
		}
	}()

	// Read input from user
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🟢 Type messages or commands like:")
	fmt.Println("   /join general")
	fmt.Println("   /leave general")
	fmt.Println("   hello world")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
	
		// Clear your typed line before sending
		fmt.Print("\r\033[K")

		if strings.HasPrefix(text, "/join ") {
			channel := strings.TrimSpace(strings.TrimPrefix(text, "/join "))
			lastChannel = channel
			stream.Send(&pb.MessageRequest{
				Sender:  username,
				Channel: channel,
				Type:    pb.RequestType_JOIN,
			})
			continue
		}

		if strings.HasPrefix(text, "/leave ") {
			channel := strings.TrimSpace(strings.TrimPrefix(text, "/leave "))
			stream.Send(&pb.MessageRequest{
				Sender:  username,
				Channel: channel,
				Type:    pb.RequestType_LEAVE,
			})
			if channel == lastChannel {
				lastChannel = ""
			}
			continue
		}

		if lastChannel == "" {
			fmt.Println("⚠️  Join a channel first using /join <channel>")
			continue
		}

		// Send message to last joined channel
		stream.Send(&pb.MessageRequest{
			Sender:  username,
			Channel: lastChannel,
			Message: text,
			Type:    pb.RequestType_MESSAGE,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Scanner error: %v", err)
	}
}