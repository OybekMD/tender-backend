package models

import (
	"errors"
	"time"

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

// ValidateTenderStatus ensures the status field is valid
func (tr *TenderUpdate) ValidateTenderStatus() error {
	return validation.ValidateStruct(
		tr,
		validation.Field(&tr.Status, validation.Required, validation.In("open", "closed", "awarded", "cancelled")),
	)
}

// ValidateTimeAndPrice ensures the deadline is in the future and the budget is non-negative
func (tc *TenderCreate) ValidateTimeAndPrice() error {
	return validation.ValidateStruct(
		tc,
		validation.Field(&tc.Deadline, validation.Required, validation.By(validateDeadline)),
		// validation.Field(&tc.Budget, validation.Required, validation.Min(0)),
	)
}

// validateDeadline checks if the deadline is a valid future time
func validateDeadline(value interface{}) error {
	deadlineStr, ok := value.(string)
	if !ok {
		return errors.New("invalid deadline format")
	}

	deadline, err := time.Parse(time.RFC3339, deadlineStr)
	if err != nil {
		return errors.New("deadline must be in RFC3339 format (e.g., 2024-11-17T10:00:00Z)")
	}

	if time.Now().After(deadline) {
		return errors.New("deadline must be a future time")
	}

	return nil
}
