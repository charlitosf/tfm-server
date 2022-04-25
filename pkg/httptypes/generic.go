package httptypes

// Generic error type struct
type Error struct {
	Message string `json:"message"`
}

// Generic error response type struct
type GenericErrorResponse struct {
	Error *Error `json:"error"`
}

// Generic empty response type struct
type GenericEmptyResponse struct {
}
