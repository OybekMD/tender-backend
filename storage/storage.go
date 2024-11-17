package storage

import (
	"tender/storage/postgres"
	"tender/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Auth() repo.AuthStorageI
	User() repo.UserStorageI
	Tender() repo.TenderStorageI
	// Bid() repo.BidStorageI
	// Notification() repo.NotificationStorageI
}

type storagePg struct {
	authRepo   repo.AuthStorageI
	userRepo   repo.UserStorageI
	tenderRepo repo.TenderStorageI
	// bidRepo          repo.UserStorageI
	// notificationRepo repo.UserStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		authRepo:   postgres.NewAuth(db),
		userRepo:   postgres.NewUser(db),
		tenderRepo: postgres.NewTender(db),
		// bidRepo:          postgres.NewUser(db),
		// notificationRepo: postgres.NewUser(db),
	}
}

func (s *storagePg) Auth() repo.AuthStorageI {
	return s.authRepo
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Tender() repo.TenderStorageI {
	return s.tenderRepo
}
