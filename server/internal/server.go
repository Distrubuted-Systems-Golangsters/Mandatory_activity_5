package internal

import (
	pb "auction/grpc"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAuctionServiceServer
	id    string
	port  string
	state *State
}

var mu sync.Mutex

var nodePorts = []string{"8080", "8081", "8082"}
var connections = make(map[string]pb.AuctionServiceClient)

var auctionTimeLength = 30
var auction *Auction

func NewServer(id string, port string) {
	server := &Server{
		id:    id,
		port:  port,
		state: NewState(port),
	}

	go server.startListening()
	go server.startHealthChecks()
}

func (s *Server) startListening() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening on %v\n", lis.Addr())

	grpcServer := grpc.NewServer()
	pb.RegisterAuctionServiceServer(grpcServer, s)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}