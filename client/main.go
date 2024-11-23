package main

import (
	"auction/client/bidder/config"
	"auction/client/bidder/core"
	"log"
)

func main() {
	cfg := config.New()
	
	auctionClient, err := core.NewAuctionClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create auction client: %v", err)
	}

	auctionClient.Start()
}