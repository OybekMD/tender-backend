package repo

import (
	"context"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"user_id"`
	Password string `json:"password"`
	Role     string `json:"role"` // client or contractor
	Email    string `json:"email"`
}

type UserStorageI interface {
	Create(ctx context.Context, user *User) (*User, error)
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context, page, limit uint64) ([]*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id string) (error)
	CheckField(ctx context.Context, field, value string) (bool, error)
}
