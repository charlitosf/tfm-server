package services

import "charlitosf/tfm-server/internal/dataaccess"

// Create file in the database
// Given propietary user, filename and file data
// Return error
func CreateFile(propietaryUser, filename string, data string) error {
	return dataaccess.CreateFile(propietaryUser, filename, data)
}

// Get a file from the database
// Given propietary user and filename
// Return file data, error
func GetFile(propietaryUser, filename string) (string, error) {
	return dataaccess.GetFile(propietaryUser, filename)
}
