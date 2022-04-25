package services

import "charlitosf/tfm-server/internal/dataaccess"

// Remove user from the database using the dataaccess functions
// Given username
// Return an error
func DeleteUser(username, token string) error {
	// Check if user exists
	_, err := dataaccess.GetUser(username)
	if err != nil {
		return err
	}

	// Logout
	err = Logout(token)
	if err != nil {
		return err
	}

	// Delete user
	return dataaccess.DeleteUser(username)
}
