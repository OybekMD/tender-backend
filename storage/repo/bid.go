package repo

import (
	"context"
)

type SubmitBidRequest struct {
	BidID        int    `json:"id"`
	TenderID     int    `json:"tender_id"`
	ContractorID string `json:"contractor_id"`
	Price        int    `json:"price"`
	DeliveryTime int `json:"delivery_time"`
	Comments     string `json:"comments"`
	BidStatus    string `json:"status"`
}

type BidStorageI interface {
	SubmitBid(ctx context.Context, req *SubmitBidRequest) (*SubmitBidRequest, error)
}
