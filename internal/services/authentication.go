package services

import (
	"bytes"
	"charlitosf/tfm-server/internal/crypt"
	"charlitosf/tfm-server/internal/dataaccess"
	"charlitosf/tfm-server/internal/jwt"
	"errors"
)

// Signup
// Given a whole user, creates it
func Signup(username, password, name, email, pubKey, privKey string) error {
	// Check if the user already exists
	_, err := dataaccess.GetUser(username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash the password using argon2
	salt := crypt.GenerateSalt()
	hashedPassword := crypt.PBKDF([]byte(password), salt)

	// Create the user
	return dataaccess.CreateUser(username, name, email, pubKey, privKey, hashedPassword, salt)
}

// Login
// Given a username and password, returns a token and the user metadata
func Login(username, password string) (*string, *dataaccess.User, error) {
	// Get the password and the salt
	pass, salt, err := dataaccess.GetUserPasswordAndSalt(username)
	if err != nil {
		return nil, nil, err
	}
	// Calculate the hash given the passed password and the salt
	hashedPassword := crypt.PBKDF([]byte(password), salt)
	// Check if passwords are equal
	if !bytes.Equal(hashedPassword, pass) {
		return nil, nil, errors.New("wrong credentials")
	}

	// Get the user
	user, err := dataaccess.GetUser(username)
	if err != nil {
		return nil, nil, err
	}

	// Generate token
	token, err := jwt.GenerateJWT(user.Username)
	if err != nil {
		return nil, nil, err
	}

	// Return the token and the user
	return &token, user, nil
}

// Logout
// Given a token, invalidates it
func Logout(token string) error {
	return jwt.RevokeToken(token)
}
