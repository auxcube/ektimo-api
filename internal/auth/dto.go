package auth

type CredentialsDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthTokenDto struct {
	Token string `json:"token"`
}
