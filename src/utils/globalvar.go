package utils

var JwtKey = []byte("authentication")

var (
	Users         = map[string]string{} // email -> hashed password
	RevokedTokens = map[string]bool{}   // token string -> revoked status
)
