package httptypes

// Generic error type struct
type Error struct {
	Message string `json:"message"`
}

// Credentials type struct
type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login request type struct
type LoginRequest struct {
	Credentials
}

// User metadata type struct
type UserMetadata struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	PubKey  string `json:"publicKey" binding:"required"`
	PrivKey string `json:"privateKey" binding:"required"`
}

// Login response type struct
type LoginResponse struct {
	User  *UserMetadata `json:"user,omitempty"`
	Token string        `json:"token,omitempty"`
	Error *Error        `json:"error,omitempty"`
}

// Signup request type struct
type SignupRequest struct {
	Credentials
	UserMetadata
}

// Signup response type struct
type SignupResponse struct {
	Error *Error `json:"error,omitempty"`
}

// Logout response type struct
type LogoutResponse struct {
	Error *Error `json:"error,omitempty"`
}
