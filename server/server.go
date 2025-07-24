package server

import (
    "context"
    "log"
    "sync"
    "time"

    pb "github.com/pgibb96/MessageApp/proto"
)

type ChatServer struct {
    pb.UnimplementedChatServiceServer
    mu       sync.Mutex
    clients  map[string]pb.ChatService_ChatStreamServer
}

func NewServer() *ChatServer {
    return &ChatServer{
        clients: make(map[string]pb.ChatService_ChatStreamServer),
    }
}

// SendMessage is a simple unary RPC (optional)
func (s *ChatServer) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
    log.Printf("[Unary] %s: %s", req.Sender, req.Message)

    return &pb.MessageResponse{
        Sender:    req.Sender,
        Message:   req.Message,
        Timestamp: time.Now().Unix(),
    }, nil
}

// ChatStream handles bidirectional chat
func (s *ChatServer) ChatStream(stream pb.ChatService_ChatStreamServer) error {
    // Read first message to register client
    first, err := stream.Recv()
    if err != nil {
        return err
    }

    sender := first.Sender
    log.Printf("Client joined: %s", sender)

    s.mu.Lock()
    s.clients[sender] = stream
    s.mu.Unlock()

    // Broadcast the first message
    s.broadcast(sender, first.Message)

    // Read incoming messages
    for {
        msg, err := stream.Recv()
        if err != nil {
            log.Printf("Client %s disconnected: %v", sender, err)
            s.mu.Lock()
            delete(s.clients, sender)
            s.mu.Unlock()
            return err
        }

        log.Printf("[%s]: %s", msg.Sender, msg.Message)
        s.broadcast(msg.Sender, msg.Message)
    }
}

func (s *ChatServer) broadcast(sender, message string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    for name, client := range s.clients {
        if name == sender {
            continue
        }

        err := client.Send(&pb.MessageResponse{
            Sender:    sender,
            Message:   message,
            Timestamp: time.Now().Unix(),
        })

        if err != nil {
            log.Printf("Failed to send to %s: %v", name, err)
        }
    }
}