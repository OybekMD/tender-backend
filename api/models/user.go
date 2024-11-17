package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	UserRequestById struct {
		UserId string `json:"user_id"`
	}

	UserRequestByUsername struct {
		Username string `json:"username"`
	}

	UserResponse struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Username     string `json:"username"`
		Bio          string `json:"bio"`
		BirthDay     string `json:"birth_day"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		Avatar       string `json:"avatar"`
		Coint        int64  `json:"coint"`
		Score        int64  `json:"score"`
		RefreshToken string `json:"refresh_token"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}

	UserUpdate struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Bio      string `json:"bio"`
		BirthDay string `json:"birth_day"`
		Avatar   string `json:"avatar"`
	}

	UserUpdatePassword struct {
		Id          string `json:"id"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
)

func (rm *UserUpdate) ValidateEmpity() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(
			&rm.Id, validation.Required,
		),
		validation.Field(
			&rm.Username, validation.Required,
		),
	)
}
