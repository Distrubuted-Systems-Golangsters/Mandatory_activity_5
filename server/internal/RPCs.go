package internal

import (
	pb "auction/grpc"
	"context"
	"log"
	"time"
)

func (s *Server) StartAuction(ctx context.Context, in *pb.StartAuctionRequest) (*pb.StartAuctionResponse, error) {
	mu.Lock()
	auction = NewAuction(auctionTimeLength)
	mu.Unlock()

	log.Printf("Auction started. Auction will end in %d seconds\n", auctionTimeLength)

	return &pb.StartAuctionResponse{}, nil
}

func (s *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	if auction == nil {
		return &pb.BidResponse{ Status: "Fail", Timestamp: int32(s.state.timestamp) }, nil
	}

	status := "Fail"

	mu.Lock()
	if(auction.endTime.After(time.Now()) && in.Amount > auction.highestBid) {
		auction.highestBid = in.Amount
		auction.bidderName = in.BidderName
		status = "Success"
	}
	mu.Unlock()

	s.state.incrementTimestamp()

	return &pb.BidResponse{ Status: status, Timestamp: int32(s.state.timestamp) }, nil
}

func (s *Server) Result(ctx context.Context, in *pb.ResultRequest) (*pb.ResultResponse, error) {
	if auction == nil {
		return &pb.ResultResponse{ Status: "Ended", HighestBid: 0, BidderName: "No bidder", Timestamp: int32(s.state.timestamp) }, nil
	}

	status := "Ongoing"

	if(auction.endTime.Before(time.Now())) {
		status = "Ended"
	}

	s.state.incrementTimestamp()

	return &pb.ResultResponse{ Status: status, HighestBid: auction.highestBid, BidderName: auction.bidderName, Timestamp: int32(s.state.timestamp) }, nil
}

func (s *Server) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	mu.Lock()
	if int(in.Timestamp) > s.state.timestamp {
		auction = &Auction{
			highestBid: in.Auction.HighestBid,
			bidderName: in.Auction.BidderName,
			endTime: in.Auction.EndTime.AsTime(),
		}
		s.state.timestamp = int(in.Timestamp)

		log.Printf("Updated auction state from server on port %s\n", s.port)
	}
	mu.Unlock()

	return &pb.HealthCheckResponse{}, nil
}