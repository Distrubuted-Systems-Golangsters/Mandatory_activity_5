package core

import (
	pb "auction/grpc"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func (c *AuctionClient) handleStatus() {
	message, err := c.getAuctionStatus()
	if err != nil {
		log.Printf("Failed to get auction status: %v\n", err)
		return
	}
	log.Println(message)
}

func (c *AuctionClient) handleStart() {
	c.startAuction()
}

func (c *AuctionClient) handleBid(input string) {
	bidAmount, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please type either 'start', 'status' or an integer:", err)
		return
	}

	message, err := c.placeBid(bidAmount)
	if err != nil {
		log.Printf("Failed to bid: %v\n", err)
		return
	}
	log.Println(message)
}

func (c *AuctionClient) startAuction() {
	var acks = make(chan bool)

	for _, client := range c.connections {
		go func(client pb.AuctionServiceClient) {
			_, err := client.StartAuction(context.Background(), &pb.StartAuctionRequest{})
			if err != nil {
				return
			}
			acks <- true
		}(client)
	}

	for i := 0; i < c.config.MinAcks; i++ {
		<-acks
	}

	log.Println("You started an auction")
}

func (c *AuctionClient) placeBid(amount int) (string, error) {
	var acks = make(chan bool)
	var mu sync.Mutex
	newestResponse := &pb.BidResponse{Status: "Failed", Timestamp: 0}

	for _, client := range c.connections {
		go func(client pb.AuctionServiceClient) {
			res, err := client.Bid(context.Background(), &pb.BidRequest{
				BidderName: c.bidderName,
				Amount:     int32(amount),
			})
			if err != nil {
				return
			}

			mu.Lock()
			if res.Timestamp > newestResponse.Timestamp {
				newestResponse = res
			}
			mu.Unlock()

			acks <- true
		}(client)
	}

	for i := 0; i < c.config.MinAcks; i++ {
		<-acks
	}

	if newestResponse.Status == "Success" {
		return "Bid successful. You are now the highest bidder", nil
	}
	return "Bid failed", nil
}

func (c *AuctionClient) getAuctionStatus() (string, error) {
	var reads = make(chan bool)
	var mu sync.Mutex
	newestResponse := &pb.ResultResponse{
		Status:     "Ongoing",
		HighestBid: 0,
		BidderName: "No bidder",
		Timestamp:  0,
	}

	for _, client := range c.connections {
		go func(client pb.AuctionServiceClient) {
			res, err := client.Result(context.Background(), &pb.ResultRequest{})
			if err != nil {
				return
			}

			mu.Lock()
			if res.Timestamp > newestResponse.Timestamp {
				newestResponse = res
			}
			mu.Unlock()

			reads <- true
		}(client)
	}

	for i := 0; i < c.config.MinReads; i++ {
		<-reads
	}

	if newestResponse.Status == "Ongoing" {
		return fmt.Sprintf("Auction is ongoing. Highest bid is %d by %s",
			newestResponse.HighestBid, newestResponse.BidderName), nil
	}
	return fmt.Sprintf("Auction has ended. Highest bid was %d by %s",
		newestResponse.HighestBid, newestResponse.BidderName), nil
}