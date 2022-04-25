package httptypes

// Update user request type struct
type UpdateUserRequest struct {
	Password string `json:"password" binding:"required"`
}
