package httptypes

// Create file request type struct
type CreateFileRequest struct {
	Name      string `json:"name" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// Get file response type struct
type GetFileResponse struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Update file request type struct
type UpdateFileRequest struct {
	Content   string `json:"content" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// Delete file request type struct
type DeleteFileRequest struct {
	Signature string `json:"signature" binding:"required"`
}
