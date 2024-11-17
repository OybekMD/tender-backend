package repo

import (
	"context"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       string `json:"id"`
	Username string `json:"user_id"`
	Password string `json:"password"`
	Role     string `json:"role"` // client or contractor
	Email    string `json:"email"`
	Access   string `json:"token"`
	Refresh  string `json:"refresh_token"`
}

type AuthStorageI interface {
	Login(ctx context.Context, username string) (*LoginResponse, error)
	Register(ctx context.Context, user *User) (*User, error)
}
