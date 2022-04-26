package httptypes

// Create file type struct
type CreateFile struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}
