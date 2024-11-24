package internal

import "time"

type Auction struct {
	highestBid int32
	bidderName string
	endTime    time.Time
}

func NewAuction(timeLength int) *Auction {
	return &Auction{
		highestBid: 0,
		bidderName: "No bidder",
		endTime:    time.Now().Add(time.Duration(timeLength) * time.Second),
	}
}
