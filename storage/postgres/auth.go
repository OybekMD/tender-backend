package postgres

import (
	"context"
	"log"
	"tender/storage/repo"

	"github.com/jmoiron/sqlx"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) repo.AuthStorageI {
	return &authRepo{
		db: db,
	}
}

// Login retrieves user information based on the provided email.
// It returns a LoginResponse struct containing the user's details or an error if the query fails.
func (l *authRepo) Login(ctx context.Context, username string) (*repo.LoginResponse, error) {
	// SQL query to fetch user details from the database.
	query := `
	SELECT
		u.id, 
		u.username, 
		u.password,
		u.role, 
		u.email
	FROM
		users u
	WHERE
		username = $1
	`

	// Define a variable to store the query result.
	var responseUser repo.LoginResponse

	// Execute the query and scan the result into the responseUser struct.
	err := l.db.QueryRowContext(ctx, query, username).Scan(
		&responseUser.ID,
		&responseUser.Username,
		&responseUser.Password,
		&responseUser.Role,
		&responseUser.Email,
	)

	// Handle potential errors during the query execution.
	if err != nil {
		// Log the error for debugging and tracking purposes.
		log.Printf("Error retrieving user from database: %v", err)
		return nil, err
	}

	// Successfully retrieved user details, returning the response.

	return &responseUser, nil
}

// Register creates a new user in the database and returns the user's details.
// It accepts a User struct containing the user information and returns a LoginResponse or an error.
func (s *authRepo) Register(ctx context.Context, user *repo.User) (*repo.User, error) {
	// SQL query to insert a new user into the database.
	query := `
	INSERT INTO users (id, username, password, role, email)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, username, role, email
	`

	// Define a variable to store the response after user registration.
	var responseUser repo.User

	// Execute the query with the provided user details and scan the returning data into responseUser.
	err := s.db.QueryRowContext(ctx, query,
		user.ID,
		user.Username,
		user.Password, // Assumes password is already hashed before being passed here.
		user.Role,
		user.Email,
	).Scan(
		&responseUser.ID,
		&responseUser.Username,
		&responseUser.Role,
		&responseUser.Email,
	)

	// Handle any error that occurs during query execution.
	if err != nil {
		log.Printf("Error registering new user: %v", err)
		return nil, err
	}

	// Successfully registered the user, return the response.
	return &responseUser, nil
}
