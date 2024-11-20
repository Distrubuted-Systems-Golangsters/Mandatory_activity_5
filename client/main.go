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
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var bidderName string
var serverPorts = []string{"localhost:8080", "localhost:8081", "localhost:8082"}
var connections []pb.AuctionServiceClient
var minAcks = 2
var minReads = 2

func main() {
	setServerConnections()
	startApp()
}

func setServerConnections() {
	for _, port := range serverPorts {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.NewClient(port, opts...)

		if err != nil {
			log.Fatalf("Connection failed: %v\n", err)
		}

		connections = append(connections, pb.NewAuctionServiceClient(conn))
	}
}

func startAuction() {
	var acks = make(chan bool)

	for _, client := range connections {
		go func() {
			_, err := client.StartAuction(context.Background(), &pb.StartAuctionRequest{})

			if err != nil {
				return
			}

			acks <- true
		}()
	}

	for i := 0; i < minAcks; i++ {
		<-acks
	}

	log.Println("You started an auction")
}

// If we don't get minAcks acknowledgements, the program will hang.
// We can handle M-N failures where M is the number of servers and N
// is the minumum number of akonowledgements required.
func placeBid(amount int) (string, error) {
	var acks = make(chan bool)
	var mu sync.Mutex

	var newestResponse = &pb.BidResponse{ Status: "Failed", Timestamp: 0 }

	for _, client := range connections {
		go func() {
			res, err := client.Bid(context.Background(), &pb.BidRequest{ BidderName: bidderName, Amount: int32(amount) })

			if err != nil {
				return
			}
		
			mu.Lock()
			if res.Timestamp > newestResponse.Timestamp {
				newestResponse = res
			}
			mu.Unlock()
	
			acks <- true
		}()
	}

	for i := 0; i < minAcks; i++ {
		<-acks
	}

	var message string
	if newestResponse.Status == "Success" {
		message = "Bid successful. You are now the highest bidder"
	} else {
		message = "Bid failed"
	}

	return message, nil
}

func getAuctionStatus() (string, error) {
	var reads = make(chan bool)
	var mu sync.Mutex

	var newestResponse = &pb.ResultResponse{ Status: "Ongoing", HighestBid: 0, BidderName: "No bidder", Timestamp: 0 }

	for _, client := range connections {
		go func() {
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
		}()
	}

	for i := 0; i < minReads; i++ {
		<-reads
	}

	var message string
	if newestResponse.Status == "Ongoing" {
		message = fmt.Sprintf("Auction is ongoing. Highest bid is %d by %s", newestResponse.HighestBid, newestResponse.BidderName)
	} else {
		message = fmt.Sprintf("Auction has ended. Highest bid was %d by %s", newestResponse.HighestBid, newestResponse.BidderName)
	}

	return message, nil
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

func startApp() {
	bidderName = getBidderName()

	for {
		reader := bufio.NewReader(os.Stdin)
		enteredString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console: %v\n", err)
		}
		enteredString = strings.ToLower(strings.Trim(enteredString, "\r\n"))

		if enteredString == "status" {
			message, err := getAuctionStatus()

			if err != nil {
				log.Fatalf("Failed to get auction status: %v\n", err)
				continue
			}

			log.Println(message)
			continue
		} else if enteredString == "start" {
			startAuction()
			continue
		}

		bidAmount, err := strconv.Atoi(enteredString)
		if err != nil {
			fmt.Println("Please type either 'start', 'status' or an integer:", err)
		}

		message, err := placeBid(bidAmount)

		if err != nil {
			log.Fatalf("Failed to bid: %v\n", err)
			continue
		}

		log.Println(message)
	}
}