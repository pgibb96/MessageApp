package server

import (
	"encoding/json"
	"log"
	"time"
	"fmt"

	"github.com/nats-io/nats.go"
	pb "github.com/pgibb96/MessageApp/proto"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	nc *nats.Conn
	js nats.JetStreamContext 
}

type inboundMessage struct {
	Sender    string `json:"sender"`
	Message   string `json:"message"`
	Channel   string `json:"channel"`
	Timestamp int64  `json:"timestamp"`
}

// NewServer creates a new ChatServer with a NATS connection and JetStream context
func NewServer(nc *nats.Conn, js nats.JetStreamContext) *ChatServer {
    return &ChatServer{
        nc: nc,
        js: js,
    }
}

func (s *ChatServer) ChatStream(stream pb.ChatService_ChatStreamServer) error {
	ctx := stream.Context()
	subs := make(map[string]*nats.Subscription)
	username := ""

	// Ensure cleanup on exit
	defer func() {
		for _, sub := range subs {
			sub.Unsubscribe()
		}
	}()

	for {
		req, err := stream.Recv()
		if err != nil {
			return nil // client disconnected
		}

		if username == "" {
			username = req.Sender
		}

		switch req.Type {
		case pb.RequestType_JOIN:
			channel := req.Channel
			subject := "chat." + channel
			consumerName := fmt.Sprintf("%s-consumer-%s", username, channel)

			if _, exists := subs[channel]; exists {
				continue
			}

			sub, err := s.js.Subscribe(subject, func(msg *nats.Msg) {
				var data inboundMessage
				if err := json.Unmarshal(msg.Data, &data); err != nil {
					log.Printf("Unmarshal error: %v", err)
					return
				}

				err = stream.Send(&pb.MessageResponse{
					Sender:    data.Sender,
					Message:   data.Message,
					Timestamp: data.Timestamp,
					Channel:   data.Channel,
				})
				if err != nil {
					log.Printf("gRPC send error: %v", err)
				} else {
					msg.Ack() // ACK after successful send
				}
			}, nats.Durable(consumerName),
				nats.ManualAck(),
				nats.AckExplicit(),
				nats.BindStream("CHAT"),
			)

			if err != nil {
				log.Printf("Error subscribing to channel %s: %v", channel, err)
				continue
			}

			subs[channel] = sub

			// Cleanup on context cancel
			go func(ch string, s *nats.Subscription) {
				<-ctx.Done()
				s.Unsubscribe()
				log.Printf("Unsubscribed %s from %s", username, ch)
			}(channel, sub)

		case pb.RequestType_LEAVE:
			channel := req.Channel
			if sub, ok := subs[channel]; ok {
				sub.Unsubscribe()
				delete(subs, channel)
			}

		case pb.RequestType_MESSAGE:
			event := inboundMessage{
				Sender:    req.Sender,
				Message:   req.Message,
				Channel:   req.Channel,
				Timestamp: time.Now().Unix(),
			}
		
			data, err := json.Marshal(event)
			if err != nil {
				log.Printf("Marshal error: %v", err)
				continue
			}
		
			subject := "chat." + req.Channel
		
			// Optionally add context-aware publish
			ack, err := s.js.PublishMsgAsync(&nats.Msg{
				Subject: subject,
				Data:    data,
				Header:  nats.Header{"sender": []string{req.Sender}},
			})
			if err != nil {
				log.Printf("Publish error: %v", err)
				continue
			}
		
			select {
			case <-ack.Ok():
				log.Printf("✅ Message published to %s", subject)
			case err := <-ack.Err():
				log.Printf("❌ JetStream publish failed: %v", err)
			case <-ctx.Done():
				log.Printf("⏹️ Publish canceled due to stream close")
			}
		}
	}
}