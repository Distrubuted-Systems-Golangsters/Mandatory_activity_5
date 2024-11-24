package internal

import (
	pb "auction/grpc"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) startHealthChecks() {
	for {
		if auction == nil {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, port := range nodePorts {
			if port == s.state.port {
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
						EndTime:    timestamppb.New(auction.endTime),
					},
					Timestamp: int32(s.state.timestamp),
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
