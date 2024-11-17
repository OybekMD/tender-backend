package repo

// import (
// 	"context"
// )

// type Notification struct {
// 	ID         uint   `json:"id"`
// 	UserID     uint   `json:"user_id"`
// 	Message    string `json:"message"`
// 	RelationID uint   `json:"relation_id"` // e.g., TenderID or BidID
// 	Type       string `json:"type"`        // e.g., bid_submitted, tender_awarded
// 	CreatedAt  string `json:"created_at"`
// }

// type NotificationStorageI interface {
// 	Create(ctx context.Context, user *Notification) (*Notification, error)
// 	Get(ctx context.Context, id string) (*Notification, error)
// 	GetAll(ctx context.Context, page, limit uint64) ([]*Notification, error)
// 	Update(ctx context.Context, user *Notification) (*Notification, error)
// 	Delete(ctx context.Context, id string) error
// 	CheckField(ctx context.Context, field, value string) (bool, error)
// }
