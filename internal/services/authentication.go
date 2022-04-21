package services

import "charlitosf/tfm-server/internal/dataaccess"

// Signup
// Given a username and a password, creates a new user
func Signup(username, password string) error {
	return dataaccess.CreateUser(username, password)
}
