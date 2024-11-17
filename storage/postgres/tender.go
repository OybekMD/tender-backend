package postgres

import (
	"context"
	"log"
	"tender/storage/repo"

	"github.com/jmoiron/sqlx"
)

type tenderRepo struct {
	db *sqlx.DB
}

func NewTender(db *sqlx.DB) repo.TenderStorageI {
	return &tenderRepo{
		db: db,
	}
}

// This function creates a tender in postgres
func (s *tenderRepo) Create(ctx context.Context, tender *repo.Tender) (*repo.Tender, error) {
	query := `
	INSERT INTO tenders (
		client_id,
		title,
		description,
		deadline,
		budget,
		status
	) VALUES ($1, $2, $3, $4, $5, $6) 
	RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		tender.ClientID,
		tender.Title,
		tender.Description,
		tender.Deadline,
		tender.Budget,
		tender.Status,
	).Scan(&tender.ID, &tender.CreatedAt, &tender.UpdatedAt)
	if err != nil {
		log.Println("Error creating tender in postgres method", err.Error())
		return nil, err
	}

	return tender, nil
}

// This function updates tender info in postgres
func (s *tenderRepo) Update(ctx context.Context, updatedTender *repo.Tender) (*repo.Tender, error) {
	query := `
	UPDATE tenders
	SET status = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = $2
	RETURNING created_at, updated_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		updatedTender.Status,
		updatedTender.ID,
	).Scan(&updatedTender.CreatedAt, &updatedTender.UpdatedAt)
	if err != nil {
		log.Println("Error updating tender in postgres method", err.Error())
		return nil, err
	}

	return updatedTender, nil
}

// This function deletes a tender in postgres
func (s *tenderRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM tenders
	WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting tender in postgres method", err.Error())
		return err
	}

	return nil
}

// This function gets a tender in postgres
func (s *tenderRepo) Get(ctx context.Context, id string) (*repo.Tender, error) {
	query := `
	SELECT id, client_id, title, description, deadline, budget, status, created_at, updated_at
	FROM tenders
	WHERE id = $1`

	var tender repo.Tender
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&tender.ID,
		&tender.ClientID,
		&tender.Title,
		&tender.Description,
		&tender.Deadline,
		&tender.Budget,
		&tender.Status,
		&tender.CreatedAt,
		&tender.UpdatedAt,
	)
	if err != nil {
		log.Println("Error getting tender in postgres method", err.Error())
		return nil, err
	}

	return &tender, nil
}

// This function gets all tenders in postgres
func (s *tenderRepo) List(ctx context.Context) ([]*repo.Tender, error) {
	query := `
	SELECT
		id,
		client_id,
		title,
		description,
		deadline,
		budget,
		status,
		created_at,
		updated_at
	FROM 
		tenders 
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error selecting tenders in postgres", err.Error())
		return nil, err
	}

	defer rows.Close()

	var tenders []*repo.Tender
	for rows.Next() {
		var tender repo.Tender
		err = rows.Scan(
			&tender.ID,
			&tender.ClientID,
			&tender.Title,
			&tender.Description,
			&tender.Deadline,
			&tender.Budget,
			&tender.Status,
			&tender.CreatedAt,
			&tender.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning tender in get all tenders method of postgres", err.Error())
			return nil, err
		}

		tenders = append(tenders, &tender)
	}

	return tenders, nil
}
