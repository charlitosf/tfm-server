package httptypes

// Generic error type struct
type Error struct {
	Message string `json:"message"`
}

// Login request type struct
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login response type struct
type LoginResponse struct {
	Token string `json:"token"`
	Error *Error `json:"error,omitempty"`
}

// Signup request type struct
type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Signup response type struct
type SignupResponse struct {
	Error *Error `json:"error,omitempty"`
}
