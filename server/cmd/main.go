package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/TheBromo/gochat/common/chat"
	"github.com/TheBromo/gochat/server/msg_distributor"

	"google.golang.org/grpc"
)

var (
	port   = flag.Int("port", 50051, "The server port")
	msgDis = msg_distributor.New()
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterChatServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedChatServiceServer
}

func (c *server) PollMesssages(msgserver pb.ChatService_ExchangeMesssagesServer) error {

	//handle input
	go func() {
		for {
			select {
			case <-msgserver.Context().Done():
				log.Fatalln(msgserver.Context().Err())
				return
			default:
				messages, err := msgserver.Recv()
				if err != nil {
					//TODO change this that it collect messages and then distributes them
					test := make([]pb.Message, 1)
					test[0] = *messages

					msgDis.Distribute(test)
				}
			}
		}

	}()

	//hande msg distribution
	go func() {
		for {
			inputC := make(chan []pb.Message)
			msgDis.RegisterConsumer(msgserver.Context(), inputC)
			defer msgDis.DeregisterConsumer(msgserver.Context())

			select {
			case <- msgserver.Context().Done():
				log.Fatalln(msgserver.Context().Err())
				return
			default:
				messages := <-inputC
				for i := 0; i < len(messages); i++ {
					msgserver.Send(&messages[i])
				}
			}
		}
	}()

	return nil
}