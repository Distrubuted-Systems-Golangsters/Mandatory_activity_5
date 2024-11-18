package main

import (
	pb "auction/grpc"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var bidderName string

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)

	if err != nil {
		log.Fatalf("Connection failed: %v\n", err)
	}

	client := pb.NewAuctionServiceClient(conn)
	startApp(client)
}

func placeBid(auctionName string, amount int, client pb.AuctionServiceClient) (string, error) {
	res, err := client.Bid(context.Background(), &pb.BidRequest{ AuctionName: auctionName, BidderName: bidderName, Amount: int32(amount), Timestamp: 0 })

	if err != nil {
		return "", err
	}

	var message string
	if res.Status == "Success" {
		message = fmt.Sprintf("Bid successful. You are now the highest bidder on the auction %s", auctionName)
	} else {
		message = "Bid failed"
	}

	return message, err
}

func getAuctionStatus(auctionName string, client pb.AuctionServiceClient) (string, error) {
	res, err := client.Result(context.Background(), &pb.ResultRequest{ AuctionName: auctionName, Timestamp: 0 })

	if err != nil {
		return "", err
	}

	var message string
	if res.Status == "Ongoing" {
		message = fmt.Sprintf("Auction is ongoing. Highest bid is %d by %s", res.HighestBid, res.BidderName)
	} else {
		message = fmt.Sprintf("Auction has ended. Highest bid was %d by %s", res.HighestBid, res.BidderName)
	}

	return message, err
}

func getBidderName() string {
	fmt.Printf("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	enteredString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console: %v\n", err)
	}
	enteredString = strings.Trim(enteredString, "\r\n")

	return enteredString
}

func startApp(client pb.AuctionServiceClient) {
	bidderName = getBidderName()

	for {
		reader := bufio.NewReader(os.Stdin)
		enteredString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console: %v\n", err)
		}
		enteredString = strings.ToLower(strings.Trim(enteredString, "\r\n"))

		auctionName := strings.Split(enteredString, " ")[0]
		secondEnter := strings.Split(enteredString, " ")[1]

		if secondEnter == "status" {
			message, err := getAuctionStatus(auctionName, client)

			if err != nil {
				log.Fatalf("Failed to get auction status: %v\n", err)
				continue
			}

			log.Println(message)
			continue
		}

		bid, err := strconv.Atoi(secondEnter)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
		}

		message, err := placeBid(auctionName, bid, client)

		if err != nil {
			log.Fatalf("Failed to bid: %v\n", err)
			continue
		}

		log.Println(message)
	}
}