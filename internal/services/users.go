package services

import (
	"charlitosf/tfm-server/internal/crypt"
	"charlitosf/tfm-server/internal/dataaccess"
)

// Remove user from the database using the dataaccess functions
// Validates the totp token and logs out the user upon deletion
// Given username, token, and totp token
// Return an error
func DeleteUser(username, token, totpToken string) error {
	// Check if user exists
	user, err := dataaccess.GetUser(username)
	if err != nil {
		return err
	}

	// Check if totp token is valid
	err = validateTOTP(user.TOTPinfo, totpToken)
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

// Update a user's password, checks if the totp token is valid
// and logs out the user upon update
// Given username, new password, token, and totp token
// Return an error
func UpdateUserPassword(username, newPassword, token string) error {
	// Check if user exists
	user, err := dataaccess.GetUser(username)
	if err != nil {
		return err
	}

	// Check if totp token is valid
	err = validateTOTP(user.TOTPinfo, token)
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
