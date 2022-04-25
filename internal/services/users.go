package services

import "charlitosf/tfm-server/internal/dataaccess"

// Remove user from the database using hte dataaccess functions
// Given username
// Return an error
func DeleteUser(username string) error {
	return dataaccess.DeleteUser(username)
}
