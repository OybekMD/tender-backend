package repo

import (
	"context"
)

type Tender struct {
	ID          uint   `json:"id"`
	ClientID    uint   `json:"client_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Budget      uint   `json:"budget"`
	Status      string `json:"status"` // open, closed, awarded
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TenderStorageI interface {
	Create(ctx context.Context, tender *Tender) (*Tender, error)
	Get(ctx context.Context, id string) (*Tender, error)
	List(ctx context.Context) ([]*Tender, error)
	Update(ctx context.Context, tender *Tender) (*Tender, error)
	Delete(ctx context.Context, id string) error
}
