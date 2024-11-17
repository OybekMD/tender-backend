package models

import (

	validation "github.com/go-ozzo/ozzo-validation"
)

type TenderCreate struct {
	ClientID    uint   `json:"client_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Budget      uint   `json:"budget"`
	Status      string `json:"status"` // open, closed, awarded
}

type TenderUpdate struct {
	ID     uint   `json:"id"`
	Status string `json:"status"` // open, closed, awarded
}

type TenderResponse struct {
	ID          uint   `json:"id"`
	ClientID    string `json:"client_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Budget      uint   `json:"budget"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (rm *TenderUpdate) ValidateTenderStatus() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Status, validation.Required, validation.In("open", "closed", "awarded", "cancelled")),
	)
}
