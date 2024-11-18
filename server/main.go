package main

import (
	pb "auction/grpc"
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAuctionServiceServer
}

type Auction struct {
	name string
	highestBid int32
	bidderName string
	endTime time.Time
}

var auctionTimeLength = 30
var auctions = make(map[string]*Auction)

func main() {
	go startListening()
	auctionPrompt()
}

func (s *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	status := "Fail"
	auction := auctions[in.AuctionName]
	
	if(auction.endTime.After(time.Now()) && in.Amount > auction.highestBid) {
		auction.highestBid = in.Amount
		auction.bidderName = in.BidderName
		status = "Success"
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

func startListening() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening on %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterAuctionServiceServer(s, &Server{})

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