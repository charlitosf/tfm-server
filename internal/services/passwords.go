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

// Get a password from a website
// Given a proprietary user, a website and a username
// Return a password
func GetPassword(proprietaryUser, website, username string) (string, error) {
	return dataaccess.GetPassword(proprietaryUser, website, username)
}

// Get passwords from a website
// Given a proprietary user and a website
// Return a map of usernames and passwords
func GetPasswords(proprietaryUser, website string) (map[string]string, error) {
	return dataaccess.GetPasswords(proprietaryUser, website)
}

// Get all the passwords
// Given a proprietary user and a totp token
// Return a map of websites and passwords
func GetAllPasswords(proprietaryUser, totpToken string) (map[string]map[string]string, error) {
	// Get the user
	user, err := dataaccess.GetUser(proprietaryUser)
	if err != nil {
		return nil, err
	}

	// Validate totp token
	err = validateTOTP(user.TOTPinfo, totpToken)
	if err != nil {
		return nil, err
	}

	return dataaccess.GetAllPasswords(proprietaryUser)
}

// Delete a password
// Given a proprietary user, a website and a username
// Return error
func DeletePassword(proprietaryUser, website, username string) error {
	// Check if password exists
	_, err := dataaccess.GetPassword(proprietaryUser, website, username)
	if err != nil {
		return errors.New("password does not exist")
	} else {
		return dataaccess.DeletePassword(proprietaryUser, website, username)
	}
}

// Update a password
// Given a proprietary user, a website, a username and a new password
// Return error
func UpdatePassword(proprietaryUser, website, username, newPassword string) error {
	// Check if password exists
	_, err := dataaccess.GetPassword(proprietaryUser, website, username)
	if err != nil {
		return errors.New("password does not exist")
	} else {
		return dataaccess.UpdatePassword(proprietaryUser, website, username, newPassword)
	}
}
