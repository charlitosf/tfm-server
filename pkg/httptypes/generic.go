package httptypes

// Generic error type struct
type Error struct {
	Message string `json:"message"`
}

// Generic response type struct
type GenericResponse struct {
	Error *Error `json:"error,omitempty"`
}
