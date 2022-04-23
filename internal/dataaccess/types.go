package dataaccess

// User type (no password)
type User struct {
	Username string `json:"username"`
	PubKey   string `json:"pubkey"`
	PrivKey  string `json:"privkey"`
}
