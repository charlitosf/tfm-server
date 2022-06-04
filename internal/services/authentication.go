package services

import (
	"bytes"
	"charlitosf/tfm-server/internal/crypt"
	"charlitosf/tfm-server/internal/dataaccess"
	"charlitosf/tfm-server/internal/jwt"
	"errors"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// Signup
// Given a whole user, creates it
func Signup(username, password, name, email, pubKey, privKey string) (string, error) {
	// Check if the user already exists
	_, err := dataaccess.GetUser(username)
	if err == nil {
		return "", errors.New("user already exists")
	}

	// Hash the password using argon2
	salt := crypt.GenerateSalt()
	hashedPassword := crypt.PBKDF([]byte(password), salt)

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "ISS Server",
		AccountName: email,
		Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		return "", err
	}

	url := key.URL()

	user := dataaccess.User{
		Username: username,
		PubKey:   pubKey,
		PrivKey:  privKey,
		Name:     name,
		Email:    email,
		TOTPinfo: url,
	}

	// Create the user
	err = dataaccess.CreateUser(user, hashedPassword, salt)
	if err != nil {
		return "", err
	}
	return url, nil
}

// Login
// Given a username, password, and TOTP token, returns a token and the user metadata
func Login(username, password, totpToken string) (*string, *dataaccess.User, error) {
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

	// Get the totp info
	key, err := otp.NewKeyFromURL(user.TOTPinfo)
	if err != nil {
		return nil, nil, err
	}

	// Check if the totp token is valid
	valid := totp.Validate(totpToken, key.Secret())
	if !valid {
		return nil, nil, errors.New("wrong credentials")
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
