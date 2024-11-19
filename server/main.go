package main

import (
	pb "auction/grpc"
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	pb.UnimplementedAuctionServiceServer
	id string
	port string
}

type Auction struct {
	name string
	highestBid int32
	bidderName string
	endTime time.Time
}

var auctionTimeLength = 30
var auctions = make(map[string]*Auction)

var nodeAddresses = []Server{
	{ id: "1", port: "8080" },
	{ id: "2", port: "8081" },
	{ id: "3", port: "8082" },
}
var maxProcessId = "3"
var nodeConnections = make(map[string]*grpc.ClientConn)

var isLeader = false

func NewServer(id string, port string) *Server {
	return &Server{ id: id, port: port }
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Arguments required: 'id' and 'port'")
		os.Exit(1)
	}

	id, port := args[0], args[1]

	if (id != "1" && id != "2" && id != "3") || (port != "8080" && port != "8081" && port != "8082") {
		fmt.Println("Invalid id or port")
		os.Exit(1)
	}

	if (id == "1" && port != "8080") || (id != "2" && port == "8081") || (id != "3" && port == "8082") {
		fmt.Println("Invalid id or port combination")
		os.Exit(1)
	}

	server := NewServer(id, port)
	go startListening(server)
	fmt.Println(isLeader)
	server.triggerLeaderElection()
	fmt.Println(isLeader)
	auctionPrompt()
}

func (s *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	status := "Fail"
	auction := auctions[in.AuctionName]
	
	if(auction.endTime.After(time.Now()) && in.Amount > auction.highestBid) {
		auction.highestBid = in.Amount
		auction.bidderName = in.BidderName
		status = "Success"

		// if(isLeader) {
			
		// }
	}

	return &pb.BidResponse{ Status: status, Timestamp: 0 }, nil
}

func (s *Server) Result(ctx context.Context, in *pb.ResultRequest) (*pb.ResultResponse, error) {
	status := "Ongoing"
	auction := auctions[in.AuctionName]

	if(auction.endTime.Before(time.Now())) {
		status = "Ended"
	}

	return &pb.ResultResponse{ Status: status, HighestBid: auction.highestBid, BidderName: auction.bidderName, Timestamp: 0 }, nil
}

func (s *Server) Election(ctx context.Context, in *pb.ElectionRequest) (*pb.ElectionResponse, error) {
	if(s.id == maxProcessId) {
		s.announceLeadership()
		return &pb.ElectionResponse{}, nil
	}
	
	s.triggerLeaderElection()
	
	return &pb.ElectionResponse{}, nil
}

func (s *Server) BroadcastLeader(ctx context.Context, in *pb.BroadcastLeaderRequest) (*pb.BroadcastLeaderResponse, error) {
	return &pb.BroadcastLeaderResponse{}, nil
}

func startListening(server *Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", server.port))
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening on %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterAuctionServiceServer(s, server)
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func startAuction() {
	reader := bufio.NewReader(os.Stdin)
	enteredString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console: %v\n", err)
	}
	enteredString = strings.ToLower(strings.Trim(enteredString, "\r\n"))

	auctions[enteredString] = &Auction{ 
		name: enteredString, 
		highestBid: 0, 
		bidderName: "No bidder",  
		endTime: time.Now().Add(time.Duration(auctionTimeLength) * time.Second),
	}

	log.Printf("Auction %v started. Auction will end in %d seconds\n", enteredString, auctionTimeLength)
}

func auctionPrompt() {
	for {
		startAuction()
	}
}

func getConnection(id string, port string) (*grpc.ClientConn, error) {
	if nodeConnections[id] != nil {
		return nodeConnections[id], nil
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(port, opts...)
	if err != nil {
		log.Fatalf("Connection to %s failed: %v\n", port, err)
		return nil, err
	}

	nodeConnections[id] = conn

	return nodeConnections[id], nil
}

func (s *Server) sendElection(port string) bool {
	conn, err := getConnection(s.id, port)

	if err != nil {
		log.Printf("Failed to connect to %s: %v", port, err)
		return false
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 1000)
	// defer cancel()

	client := pb.NewAuctionServiceClient(conn)
	_, err = client.Election(context.Background(), &pb.ElectionRequest{ Id: s.id })
	if err != nil {
		log.Printf("Election request to %s", port)
		return false
	}

	log.Printf("Election request to %s succeeded", port)
	return true
}

func (s *Server) triggerLeaderElection() {
	timeout := time.Duration(5 * time.Second)
	responseRecieved := make(chan bool)

	for _, node := range nodeAddresses {
		if node.id > s.id {
			go func(port string) {
				if s.sendElection(port) {
					responseRecieved <- true
				}
			}(node.port)
		}
	}

	select {
		case <-responseRecieved: // If any response is received
			log.Printf("Node %s received response so it is not the leader", s.id)
		case <-time.After(timeout): // If timeout occurs
			log.Printf("Node %s received no response; declaring itself leader", s.id)
			isLeader = true
			s.announceLeadership()
	}
}

func (s *Server) announceLeadership() {
	for _, node := range nodeAddresses {
		if node.id < s.id {
			go func(port string) {
				conn, err := getConnection(s.id, port)

				if err != nil {
					log.Printf("Failed to connect to %s: %v", port, err)
					return
				}

				client := pb.NewAuctionServiceClient(conn)
				client.BroadcastLeader(context.Background(), &pb.BroadcastLeaderRequest{ Leader: s.id })
			}(node.port)
		}
	}
}