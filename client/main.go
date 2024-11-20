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
var minAcks = 1
var minReads = 1

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

// If we don't get minAcks acknowledgements, the program will hang.
// We can handle M-N failures where M is the number of servers and N
// is the minumum number of akonowledgements required.
func placeBid(amount int) (string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var newestResponse = &pb.BidResponse{ Status: "Failed", Timestamp: 0 }

	wg.Add(minAcks)
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
	
			wg.Done()
		}()
	}

	wg.Wait()

	var message string
	if newestResponse.Status == "Success" {
		message = "Bid successful. You are now the highest bidder"
	} else {
		message = "Bid failed"
	}

	return message, nil
}

func getAuctionStatus() (string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var newestResponse = &pb.ResultResponse{ Status: "Ongoing", HighestBid: 0, BidderName: "No bidder", Timestamp: 0 }

	wg.Add(minReads)
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
	
			wg.Done()
		}()
	}

	wg.Wait()

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
		}

		bidAmount, err := strconv.Atoi(enteredString)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
		}

		message, err := placeBid(bidAmount)

		if err != nil {
			log.Fatalf("Failed to bid: %v\n", err)
			continue
		}

		log.Println(message)
	}
}