package repo

import (
	"context"
)

type SubmitBidRequest struct {
	BidID        int    `json:"bid_id"`
	TenderID     int    `json:"tender_id"`
	ContractorID string `json:"contractor_id"`
	Price        int    `json:"price"`
	DeliveryTime string `json:"delivery_time"`
	Comments     string `json:"comments"`
	BidStatus    string `json:"status"`
}

type BidStorageI interface {
	SubmitBid(ctx context.Context, req *SubmitBidRequest) (*SubmitBidRequest, error)
}
