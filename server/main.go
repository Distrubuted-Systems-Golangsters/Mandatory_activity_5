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
)

type Server struct {
	pb.UnimplementedAuctionServiceServer
	id string
	port string
}

type Auction struct {
	highestBid int32
	bidderName string
	endTime time.Time
}

var auctionTimeLength = 30
var auction *Auction
var timestamp = 0

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
	auctionPrompt()
}

func (s *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	status := "Fail"
	
	if(auction.endTime.After(time.Now()) && in.Amount > auction.highestBid) {
		auction.highestBid = in.Amount
		auction.bidderName = in.BidderName
		status = "Success"
	}

	timestamp++

	return &pb.BidResponse{ Status: status, Timestamp: int32(timestamp) }, nil
}

func (s *Server) Result(ctx context.Context, in *pb.ResultRequest) (*pb.ResultResponse, error) {
	status := "Ongoing"

	if(auction.endTime.Before(time.Now())) {
		status = "Ended"
	}

	timestamp++

	return &pb.ResultResponse{ Status: status, HighestBid: auction.highestBid, BidderName: auction.bidderName, Timestamp: int32(timestamp) }, nil
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

	if enteredString != "start" {
		return
	}

	auction = &Auction{ 
		highestBid: 0, 
		bidderName: "No bidder",  
		endTime: time.Now().Add(time.Duration(auctionTimeLength) * time.Second),
	}

	log.Printf("Auction started. Auction will end in %d seconds\n", auctionTimeLength)
}

func auctionPrompt() {
	for {
		startAuction()
	}
}