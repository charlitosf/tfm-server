package services

import (
	"charlitosf/tfm-server/internal/dataaccess"
	"errors"
)

// Create file in the database
// Given propietary user, filename and file data
// Return error
func CreateFile(propietaryUser, filename string, data string) error {
	// Check if a file with that name already exists
	_, err := dataaccess.GetFile(propietaryUser, filename)
	if err == nil {
		return errors.New("file already exists")
	}
	return dataaccess.CreateFile(propietaryUser, filename, data)
}

// Get a file from the database
// Given propietary user and filename
// Return file data, error
func GetFile(propietaryUser, filename string) (string, error) {
	return dataaccess.GetFile(propietaryUser, filename)
}

// Delete a file from the database
// Given propietary user and filename
// Return error
func DeleteFile(propietaryUser, filename string) error {
	// Check if file exists
	_, err := dataaccess.GetFile(propietaryUser, filename)
	if err != nil {
		return errors.New("file not found")
	}
	return dataaccess.DeleteFile(propietaryUser, filename)
}

// Update a file in the database
// Given propietary user, filename and file data
// Return error
func UpdateFile(propietaryUser, filename string, data string) error {
	// Check if file exists
	_, err := dataaccess.GetFile(propietaryUser, filename)
	if err != nil {
		return errors.New("file not found")
	}
	return dataaccess.UpdateFile(propietaryUser, filename, data)
}
