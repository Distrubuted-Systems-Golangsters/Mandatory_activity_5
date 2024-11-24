package core

import (
	"auction/client/bidder/config"
	pb "auction/grpc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

		parts := strings.Fields(enteredString)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		switch command {
		case "status":
			c.handleStatus()
		case "start":
			if len(parts) != 2 {
				fmt.Println("Please provide duration of auction in seconds")
				continue
			}
			duration, err := strconv.Atoi(parts[1])
			if err != nil || duration <= 0 {
				fmt.Println("Please provide a valid positive integer for the duration.")
				continue
			}
			c.handleStart(duration)
		default:
			c.handleBid(enteredString)
		}
	}
}
