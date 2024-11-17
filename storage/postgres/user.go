package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"tender/storage/repo"
)

type userRepo struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

// Create a new user in the database
func (u *userRepo) Create(ctx context.Context, user *repo.User) (*repo.User, error) {
	query := `
	INSERT INTO users (username, password, role, email) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id`

	err := u.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Role,
		user.Email,
	).Scan(&user.ID)
	if err != nil {
		log.Println("Error creating user in postgres method", err.Error())
		return nil, err
	}
	return user, nil
}

// Get a user by ID
func (u *userRepo) Get(ctx context.Context, id string) (*repo.User, error) {
	query := `
	SELECT id, username, password, role, email 
	FROM users 
	WHERE id = $1`

	var user repo.User
	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Email,
	)
	if err != nil {
		log.Println("Error getting user in postgres method", err.Error())
		return nil, err
	}
	return &user, nil
}

// Get all users with pagination
func (u *userRepo) GetAll(ctx context.Context, page, limit uint64) ([]*repo.User, error) {
	query := `
	SELECT id, username, password, role, email 
	FROM users 
	LIMIT $1 OFFSET $2`

	offset := limit * (page - 1)
	rows, err := u.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error selecting users in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []*repo.User
	for rows.Next() {
		var user repo.User
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Email,
		)
		if err != nil {
			log.Println("Error scanning user in get all users method of postgres", err.Error())
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// Update user information
func (u *userRepo) Update(ctx context.Context, user *repo.User) (*repo.User, error) {
	query := `
	UPDATE users 
	SET username = $1, password = $2, role = $3, email = $4 
	WHERE id = $5 
	RETURNING id`

	err := u.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Role,
		user.Email,
		user.ID,
	).Scan(&user.ID)
	if err != nil {
		log.Println("Error updating user in postgres method", err.Error())
		return nil, err
	}
	return user, nil
}

// Delete a user by ID
func (u *userRepo) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM users 
	WHERE id = $1`

	_, err := u.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting user in postgres method", err.Error())
		return err
	}
	return nil
}

func (u *userRepo) CheckField(ctx context.Context, field, value string) (bool, error) {

	query := fmt.Sprintf(
		`SELECT count(1) 
		FROM users WHERE %s = $1`, field)

	var isExists int

	row := u.db.QueryRowContext(ctx, query, value)
	if err := row.Scan(&isExists); err != nil {
		return true, err
	}

	if isExists == 0 {
		return false, nil
	}

	return true, nil
}
