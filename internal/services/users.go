package services

import (
	"charlitosf/tfm-server/internal/crypt"
	"charlitosf/tfm-server/internal/dataaccess"
)

// Remove user from the database using the dataaccess functions
// Given username
// Return an error
func DeleteUser(username, token string) error {
	// Check if user exists
	_, err := dataaccess.GetUser(username)
	if err != nil {
		return err
	}

	// Delete user
	err = dataaccess.DeleteUser(username)
	if err != nil {
		return err
	}

	// Logout
	return Logout(token)
}

// Update a user's password
// Given username and new password
// Return an error
func UpdateUserPassword(username, newPassword, token string) error {
	// Check if user exists
	_, err := dataaccess.GetUser(username)
	if err != nil {
		return err
	}

	// Generate new salt
	salt := crypt.GenerateSalt()
	// Hash password
	hashedPassword := crypt.PBKDF([]byte(newPassword), salt)

	// Update user
	err = dataaccess.UpdateUserPassword(username, hashedPassword, salt)
	if err != nil {
		return err
	}

	// Logout
	return Logout(token)
}
