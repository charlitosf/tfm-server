package httptypes

// Create password request type struct
type CreatePasswordRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Create password response type struct
type CreatePasswordResponse struct {
	Error *Error `json:"error,omitempty"`
}

// Get passwords response type struct
type GetPasswordsResponse struct {
	Error *Error `json:"error,omitempty"`
}
