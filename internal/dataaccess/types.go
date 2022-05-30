package dataaccess

// User type (no password)
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PubKey   string `json:"pubkey"`
	PrivKey  string `json:"privkey"`
	TOTPinfo string `json:"totp"`
}
