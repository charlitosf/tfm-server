package services

import "charlitosf/tfm-server/internal/dataaccess"

// Create a new password
// Given a proprietary user, a website, a username and a password
// Return error
func CreatePassword(user, website, username, password string) error {
	return dataaccess.CreatePassword(user, website, username, password)
}
