package core

import (
	"auction/client/bidder/config"
	pb "auction/grpc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type AuctionClient struct {
	connections []pb.AuctionServiceClient
	bidderName  string
	config      *config.Config
}

func NewAuctionClient(cfg *config.Config) (*AuctionClient, error) {
	client := &AuctionClient{
		config: cfg,
	}

	if err := client.setupConnections(); err != nil {
		return nil, fmt.Errorf("failed to setup connections: %w", err)
	}

	client.bidderName = client.getBidderName()
	return client, nil
}

func (c *AuctionClient) Start() {
	for {
		reader := bufio.NewReader(os.Stdin)
		enteredString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console: %v\n", err)
		}
		enteredString = strings.ToLower(strings.Trim(enteredString, "\r\n"))

		switch enteredString {
			case "status":
				c.handleStatus()
			case "start":
				c.handleStart()
			default:
				c.handleBid(enteredString)
		}
	}
}