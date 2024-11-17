package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"tender/storage/repo"
)

type bidRepo struct {
	db *sqlx.DB
}

func NewBidRepo(db *sqlx.DB) repo.BidStorageI {
	return &bidRepo{db: db}
}

func (b *bidRepo) SubmitBid(ctx context.Context, bid *repo.SubmitBidRequest) (*repo.SubmitBidRequest, error) {
	query := `
	INSERT INTO bids (tender_id, contractor_id, price, delivery_time, comments, status)
	VALUES ($1, $2, $3, $4, $5, 'pending')
	RETURNING id, tender_id, contractor_id, price, delivery_time, comments, status;
	`

	var response repo.SubmitBidRequest

	err := b.db.QueryRowContext(ctx, query,
		bid.TenderID,
		bid.ContractorID,
		bid.Price,
		bid.DeliveryTime,
		bid.Comments,
	).Scan(
		&response.BidID,
		&response.TenderID,
		&response.ContractorID,
		&response.Price,
		&response.DeliveryTime,
		&response.Comments,
		&response.BidStatus,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
