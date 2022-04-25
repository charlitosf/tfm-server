package httptypes

// Delete user response type struct
type DeleteUserResponse struct {
	Error *Error `json:"error,omitempty"`
}

// Update user request type struct
type UpdateUserRequest struct {
	Password string `json:"password" binding:"required"`
}

// Update user response type struct
type UpdateUserResponse struct {
	Error *Error `json:"error,omitempty"`
}
