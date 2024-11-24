package core

import (
	pb "auction/grpc"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *AuctionClient) setupConnections() error {
	for _, port := range c.config.ServerPorts {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.NewClient(port, opts...)
		if err != nil {
			return fmt.Errorf("failed to connect to %s: %w", port, err)
		}

		c.connections = append(c.connections, pb.NewAuctionServiceClient(conn))
	}

	return nil
}

func (c *AuctionClient) getBidderName() string {
	fmt.Printf("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	enteredString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console: %v\n", err)
	}
	return strings.Trim(enteredString, "\r\n")
}
