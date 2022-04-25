package services

import (
	"charlitosf/tfm-server/internal/dataaccess"
	"errors"
)

// Create a new password
// Given a proprietary user, a website, a username and a password
// Return error
func CreatePassword(proprietaryUser, website, username, password string) error {
	// Check if password already exists
	_, err := dataaccess.GetPassword(proprietaryUser, website, username)
	if err != nil {
		// Create password
		return dataaccess.CreatePassword(proprietaryUser, website, username, password)
	} else {
		return errors.New("password already exists")
	}
}
