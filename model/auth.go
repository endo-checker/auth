package model

type UserInfo struct {
	Sub       string `json:"sub"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Picture   string `json:"picture"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
}

type SignIn struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int32  `json:"expires_in"`
	IdToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
}
