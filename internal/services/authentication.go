package services

import (
	"charlitosf/tfm-server/internal/dataaccess"
	"errors"
)

// Signup
// Given a username and a password, creates a new user
func Signup(username, password, name, email, pubKey, privKey string) error {
	// Check if the user already exists
	_, err := dataaccess.GetUser(username)
	if err == nil {
		return errors.New("user already exists")
	}
	return dataaccess.CreateUser(username, password, name, email, pubKey, privKey)
}
