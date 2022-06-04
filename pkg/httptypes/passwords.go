package httptypes

// Create password request type struct
type CreatePasswordRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// Get password response type struct
type GetPasswordResponse struct {
	Error     *Error `json:"error,omitempty"`
	Password  string `json:"password,omitempty"`
	Signature string `json:"signature,omitempty"`
}

// Update password request type struct
type UpdatePasswordRequest struct {
	Password  string `json:"password" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

// Delete password request type struct
type DeletePasswordRequest struct {
	Signature string `json:"signature" binding:"required"`
}
