package main

import (
	pb "auction/grpc"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedAuctionServiceServer
	id string
	port string
}

type Auction struct {
	highestBid int32
	bidderName string
	endTime time.Time
}

type State struct {
	timestamp int
	port string
	mu sync.Mutex
}

var mu sync.Mutex

var nodePorts = []string{"8080", "8081", "8082"}
var connections = make(map[string]pb.AuctionServiceClient)

var auctionTimeLength = 30
var auction *Auction

var state *State

func NewServer(id string, port string) *Server {
	return &Server{ id: id, port: port }
}

func NewState(port string) *State {
	return &State{ timestamp: 0, port: port }
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Arguments required: 'id' and 'port'")
		os.Exit(1)
	}

	id, port := args[0], args[1]

	if (id != "1" && id != "2" && id != "3") || (port != "8080" && port != "8081" && port != "8082") {
		fmt.Println("Invalid id or port")
		os.Exit(1)
	}

	if (id == "1" && port != "8080") || (id != "2" && port == "8081") || (id != "3" && port == "8082") {
		fmt.Println("Invalid id or port combination")
		os.Exit(1)
	}

	server := NewServer(id, port)
	state = NewState(port)
	go startListening(server)
	go startHealthChecks()

	//blocker
	bl := make(chan bool)
	<-bl
}

func (s *Server) StartAuction(ctx context.Context, in *pb.StartAuctionRequest) (*pb.StartAuctionResponse, error) {
	mu.Lock()
	auction = &Auction{
		highestBid: 0,
		bidderName: "No bidder",
		endTime: time.Now().Add(time.Duration(auctionTimeLength) * time.Second),
	}
	mu.Unlock()

	log.Printf("Auction started. Auction will end in %d seconds\n", auctionTimeLength)

	return &pb.StartAuctionResponse{}, nil
}

func (s *Server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	if auction == nil {
		return &pb.BidResponse{ Status: "Fail", Timestamp: int32(state.timestamp) }, nil
	}

	status := "Fail"

	mu.Lock()
	if(auction.endTime.After(time.Now()) && in.Amount > auction.highestBid) {
		auction.highestBid = in.Amount
		auction.bidderName = in.BidderName
		status = "Success"
	}
	mu.Unlock()

	state.incrementTimestamp()

	return &pb.BidResponse{ Status: status, Timestamp: int32(state.timestamp) }, nil
}

func (s *Server) Result(ctx context.Context, in *pb.ResultRequest) (*pb.ResultResponse, error) {
	if auction == nil {
		return &pb.ResultResponse{ Status: "Ended", HighestBid: 0, BidderName: "No bidder", Timestamp: int32(state.timestamp) }, nil
	}

	status := "Ongoing"

	if(auction.endTime.Before(time.Now())) {
		status = "Ended"
	}

	state.incrementTimestamp()

	return &pb.ResultResponse{ Status: status, HighestBid: auction.highestBid, BidderName: auction.bidderName, Timestamp: int32(state.timestamp) }, nil
}

func (s *Server) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	mu.Lock()
	if int(in.Timestamp) > state.timestamp {
		auction = &Auction{
			highestBid: in.Auction.HighestBid,
			bidderName: in.Auction.BidderName,
			endTime: in.Auction.EndTime.AsTime(),
		}
		state.timestamp = int(in.Timestamp)

		log.Printf("Updated auction state from server on port %s\n", s.port)
	}
	mu.Unlock()

	return &pb.HealthCheckResponse{}, nil
}

func startListening(server *Server) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", server.port))
	if err != nil {
		panic(err)
	}

	log.Printf("Server listening on %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterAuctionServiceServer(s, server)
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func startHealthChecks() {
	for {
		if auction == nil {
            time.Sleep(5 * time.Second)
            continue
        }

		for _, port := range nodePorts {
			if port == state.port {
				continue
			}

			go func(port string) {
				client, err := getConnection(port)

				if err != nil {
					log.Printf("Failed to connect to server on port %s\n", port)

					mu.Lock()
					delete(connections, port)
					mu.Unlock()

					return
				}

				_, err = client.HealthCheck(context.Background(), &pb.HealthCheckRequest{ 
					Auction: &pb.Auction{ 
						HighestBid: auction.highestBid, 
						BidderName: auction.bidderName, 
						EndTime: timestamppb.New(auction.endTime),
					},
					Timestamp: int32(state.timestamp),  
				})

				if err != nil {
					log.Printf("Failed to send health check to server on port %s\n", port)
					mu.Lock()
                    delete(connections, port)
                    mu.Unlock()
				}
			}(port)
		}

		time.Sleep(5 * time.Second)
	}
}

func getConnection(port string) (pb.AuctionServiceClient, error) {
	if connections[port] != nil {
		return connections[port], nil
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%s", port), opts...)

	if err != nil {
		return nil, err
	}

	mu.Lock()
	connections[port] = pb.NewAuctionServiceClient(conn)
	mu.Unlock()

	return connections[port], nil
}

func (st *State) incrementTimestamp() {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.timestamp++
}