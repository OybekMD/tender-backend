package hashing

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// This function hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println(bytes)
	return string(bytes), err
}

// This function chech hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
