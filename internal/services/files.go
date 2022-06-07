package services

import (
	"charlitosf/tfm-server/internal/dataaccess"
	"errors"
)

// Create file in the database
// Given propietary user, filename, file data, and its signature
// Return error
func CreateFile(propietaryUser, filename, data, signature string) error {
	// Check if a file with that name already exists
	_, err := dataaccess.GetFile(propietaryUser, filename)
	if err == nil {
		return errors.New("file already exists")
	}

	err = validateSignature(propietaryUser, data, signature)
	if err != nil {
		return err
	}

	return dataaccess.CreateFile(propietaryUser, filename, data, signature)
}

// Get a file from the database
// Given propietary user and filename
// Return file data, error
func GetFile(propietaryUser, filename string) (string, string, error) {
	file, err := dataaccess.GetFile(propietaryUser, filename)
	if err != nil {
		return "", "", err
	}
	return file.Contents, file.Signature, nil
}

// Delete a file from the database
// Given propietary user, filename, and its signature
// Return error
func DeleteFile(propietaryUser, filename, signature string) error {
	// Check if file exists
	_, err := dataaccess.GetFile(propietaryUser, filename)
	if err != nil {
		return errors.New("file not found")
	}

	err = validateSignature(propietaryUser, filename, signature)
	if err != nil {
		return err
	}

	return dataaccess.DeleteFile(propietaryUser, filename)
}

// Update a file in the database
// Given propietary user, filename, file data, and its signature
// Return error
func UpdateFile(propietaryUser, filename, data, signature string) error {
	// Check if file exists
	_, err := dataaccess.GetFile(propietaryUser, filename)
	if err != nil {
		return errors.New("file not found")
	}

	err = validateSignature(propietaryUser, data, signature)
	if err != nil {
		return err
	}

	return dataaccess.UpdateFile(propietaryUser, filename, data, signature)
}

// Get all filenames from the database
// Given propietary user
// Return filenames, error
func GetAllFilenames(propietaryUser string) ([]map[string]string, error) {
	return dataaccess.GetAllFilenames(propietaryUser)
}

// Delete all files from the database
// Given propietary user
// Return error
func DeleteAllFiles(propietaryUser string) error {
	return dataaccess.DeleteAllFiles(propietaryUser)
}
