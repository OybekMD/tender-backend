package repo

// import (
// 	"context"
// )

// type Bid struct {
// 	ID           uint    `json:"id"`
// 	TenderID     uint    `json:"tender_id"`
// 	ContractorID uint    `json:"contractor_id"`
// 	Price        float64 `json:"price"`
// 	DeliveryTime int     `json:"delivery_time"`
// 	Comments     string  `json:"comments"`
// 	Status       string  `json:"status"` // pending, accepted, rejected
// }

// type BidStorageI interface {
// 	Create(ctx context.Context, user *Bid) (*Bid, error)
// 	Get(ctx context.Context, id string) (*Bid, error)
// 	GetAll(ctx context.Context, page, limit uint64) ([]*Bid, error)
// 	Update(ctx context.Context, user *Bid) (*Bid, error)
// 	Delete(ctx context.Context, id string) error
// 	CheckField(ctx context.Context, field, value string) (bool, error)
// }
