package services

import (
	"charlitosf/tfm-server/internal/crypt"
	"charlitosf/tfm-server/internal/dataaccess"
	"errors"
)

// Create a new password
// Given a proprietary user, a website, a username, a password, and its signature
// Return error
func CreatePassword(proprietaryUser, website, username, password, signature string) error {
	// Check if password already exists
	_, err := dataaccess.GetPassword(proprietaryUser, website, username)
	if err == nil {
		return errors.New("password already exists")
	}

	// Get user
	user, err := dataaccess.GetUser(proprietaryUser)
	if err != nil {
		return err
	}

	// Get its public key
	pubKey, err := crypt.DecodePubKey(user.PubKey)
	if err != nil {
		return err
	}

	// Decode 64 signature
	signatureBytes, err := crypt.Decode64(signature)
	if err != nil {
		return err
	}

	// Verify signature
	err = crypt.VerifyRSASignature(password, signatureBytes, pubKey)
	if err != nil {
		return err
	}

	// Create password
	return dataaccess.CreatePassword(proprietaryUser, website, username, password, signature)
}

// Get a password from a website
// Given a proprietary user, a website and a username
// Return a password
func GetPassword(proprietaryUser, website, username string) (*dataaccess.Password, error) {
	return dataaccess.GetPassword(proprietaryUser, website, username)
}

// Get passwords from a website
// Given a proprietary user and a website
// Return a map of usernames and passwords
func GetPasswords(proprietaryUser, website string) (map[string]dataaccess.Password, error) {
	return dataaccess.GetPasswords(proprietaryUser, website)
}

// Get all the passwords
// Given a proprietary user and a totp token
// Return a map of websites and passwords
func GetAllPasswords(proprietaryUser, totpToken string) (map[string]map[string]dataaccess.Password, error) {
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
// Given a proprietary user, a website, a username, and a signature
// Return error
func DeletePassword(proprietaryUser, website, username, signature string) error {
	// Check if password exists
	_, err := dataaccess.GetPassword(proprietaryUser, website, username)
	if err != nil {
		return errors.New("password does not exist")
	}

	// Get user
	user, err := dataaccess.GetUser(proprietaryUser)
	if err != nil {
		return err
	}
	// Get its public key
	pubKey, err := crypt.DecodePubKey(user.PubKey)
	if err != nil {
		return err
	}

	// Decode 64 signature
	signatureBytes, err := crypt.Decode64(signature)
	if err != nil {
		return err
	}
	// Verify signature
	err = crypt.VerifyRSASignature(website+username, signatureBytes, pubKey)
	if err != nil {
		return err
	}

	return dataaccess.DeletePassword(proprietaryUser, website, username)
}

// Update a password
// Given a proprietary user, a website, a username, a new password, and its signature
// Return error
func UpdatePassword(proprietaryUser, website, username, newPassword, signature string) error {
	// Check if password exists
	_, err := dataaccess.GetPassword(proprietaryUser, website, username)
	if err != nil {
		return errors.New("password does not exist")
	}

	// Get user
	user, err := dataaccess.GetUser(proprietaryUser)
	if err != nil {
		return err
	}

	// Get its public key
	pubKey, err := crypt.DecodePubKey(user.PubKey)
	if err != nil {
		return err
	}

	// Decode 64 signature
	signatureBytes, err := crypt.Decode64(signature)
	if err != nil {
		return err
	}

	// Verify signature
	err = crypt.VerifyRSASignature(newPassword, signatureBytes, pubKey)
	if err != nil {
		return err
	}

	return dataaccess.UpdatePassword(proprietaryUser, website, username, newPassword, signature)
}
