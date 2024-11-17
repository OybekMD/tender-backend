package models

import (
	// "regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	Register struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // client or contractor
		Email    string `json:"email"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // client or contractor
		Email    string `json:"email"`
		Token    string `json:"token"`
		Refresh  string `json:"refresh_token"`
	}

	RegisterResponse struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // client or contractor
		Email    string `json:"email"`
		Token    string `json:"token"`
		Refresh  string `json:"refresh_token"`
	}
)

func (rm *Register) ValidateEmail() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Email, validation.Required, is.Email),
	)
}

func (rm *Register) ValidateRole() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Role, validation.Required, validation.In("admin", "client", "contractor")),
	)
}

// func (rm *Register) ValidatePassword() error {
// 	return validation.ValidateStruct(
// 		rm,
// 		validation.Field(
// 			&rm.Password,
// 			validation.Required,
// 			validation.Length(8, 30),
// 			validation.Match(regexp.MustCompile("[a-z]|[1-9]")),
// 		),
// 	)
// }
